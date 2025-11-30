package server

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	apiError2 "github.com/dozro/tawny/pkg/apiError"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func generateHMAC(secret, message string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	signature := h.Sum(nil)
	return hex.EncodeToString(signature)
}

func verifyHMAC(secret, message, receivedSig string) bool {
	log.Debugf("verifyHMAC receivedSig: %s", receivedSig)
	key := []byte(secret)
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)

	receivedMAC, err := hex.DecodeString(receivedSig)
	if err != nil {
		log.Errorf("verifyHMAC failed to decode signature: %s", err)
		return false
	}

	return hmac.Equal(expectedMAC, receivedMAC)
}
func signBase64Request(c *gin.Context) {
	psk := c.Request.Header.Get(stringsHmacPsk)
	raw, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiError2.ApiError{
			HttpCode:          403,
			InternalErrorCode: apiError2.InvalidBody,
			Message:           stringsInvalidRequestBody,
			Success:           false,
		})
		return
	}

	signature := generateHMAC(psk, string(raw))

	log.Debug("signing request")
	log.Debugf("Message bytes: %s", raw)
	log.Printf("Generated signature: %s", signature)

	c.JSON(http.StatusOK, HmacBase64SignedRequest{
		Signature: signature,
		Request:   raw,
	})
}
func signRequest(c *gin.Context) {
	psk := c.Request.Header.Get(stringsHmacPsk)

	raw, err := io.ReadAll(c.Request.Body)
	if err != nil {
		render(c, http.StatusBadRequest, apiError2.ApiError{
			HttpCode:          403,
			InternalErrorCode: apiError2.InvalidBody,
			Message:           stringsInvalidRequestBody,
			Success:           false,
		})
		return
	}

	signature := generateHMAC(psk, string(raw))

	log.Debug(logStringSigningRequest)
	log.Debugf("Message bytes: %s", raw)
	log.Printf("Generated signature: %s", signature)

	signedRequest := HmacSignedRequest{
		Request:   raw,
		Signature: signature,
	}

	render(c, 200, signedRequest)
}

func verifyRequest(c *gin.Context) {
	psk := c.Request.Header.Get(stringsHmacPsk)
	isValid, _, err, code := verifyRequestInternal(c, psk, determineIfBase64(c), nil)
	if err != nil {
		render(c, code, apiError2.ApiError{
			HttpCode: code,
			Message:  err.Error(),
			Success:  false,
		})
		return
	}
	if !isValid {
		render(c, 403, apiError2.ApiError{
			HttpCode: 403,
			Message:  stringsInvalidHmacSignature,
			Data:     c.Request.Body,
			Success:  false,
		})
		return
	}
	render(c, 200, gin.H{
		"message": stringsValidHmacSignature,
		"success": true,
	})
}

func verifyAgainstServerSecret(c *gin.Context) {
	isValid, _, err, code := verifyRequestInternal(c, proxyConfig.HmacSecret, determineIfBase64(c), nil)
	if err != nil {
		render(c, code, apiError2.ApiError{
			HttpCode: code,
			Message:  err.Error(),
			Success:  false,
		})
	}
	if !isValid {
		render(c, 403, apiError2.ApiError{
			HttpCode: 403,
			Message:  stringsInvalidHmacSignature,
			Success:  false,
		})
		return
	}
	render(c, 200, gin.H{
		"message": stringsValidHmacSignature,
		"success": true,
	})
}

// verifyRequestInternal internal code for hmac request verification
func verifyRequestInternal(c *gin.Context, hmacSecret string, base64 bool, overridenReqCont *HmacBase64SignedRequest) (bool, *HmacProxyRequest, error, int) {
	var signedReq HmacSignedRequest
	var isValid bool

	if base64 {
		var signedReqBase64 HmacBase64SignedRequest
		if overridenReqCont != nil {
			log.Debug("overriding request body")
			signedReqBase64.Request = overridenReqCont.Request
			signedReqBase64.Signature = overridenReqCont.Signature
		} else {
			if c.ShouldBindJSON(&signedReqBase64) != nil {
				return false, nil, errors.New("invalid JSON"), http.StatusBadRequest
			}
		}
		log.Debug("verifying signature (base64)...")
		log.Debugf("Received message bytes: %s", signedReqBase64.Request)
		log.Debugf("Received signature: %s", signedReqBase64.Signature)

		var err error
		signedReq, err = Base64ToHmacSignedRequest(signedReqBase64)
		if err != nil {
			return false, nil, err, http.StatusBadRequest
		}

	} else {
		if c.ShouldBindJSON(&signedReq) != nil {
			return false, nil, errors.New("invalid JSON"), http.StatusBadRequest
		}

		log.Debug("verifying signature (not base64)...")
		log.Debugf("Received message bytes: %s", signedReq.Request)
		log.Debugf("Received signature: %s", signedReq.Signature)

	}

	isValid = verifyHMAC(hmacSecret, string(signedReq.Request), signedReq.Signature)

	if !isValid {
		log.Debug(stringsInvalidHmacSignature)
		log.Debugf("Signature: %s", signedReq.Signature)
		log.Debugf("Signature should be: %s", generateHMAC(hmacSecret, string(signedReq.Request)))
		return false, nil, errors.New("invalid HMAC signature"), http.StatusForbidden
	}

	log.Debug("successfully verified request")
	log.Debug("parsing request...")

	var req HmacProxyRequest
	if err := json.Unmarshal(signedReq.Request, &req); err != nil {
		log.Errorf("failed to unmarshal request: %s", err)
		return false, nil, err, http.StatusBadRequest
	}

	return true, &req, nil, http.StatusOK
}

func executeSignedRequest(c *gin.Context) {
	var overridenCont *HmacBase64SignedRequest
	if c.Request.Method == "GET" {
		request, err := url.QueryUnescape(c.Query("request"))
		log.Debugf("Request: %s, %s", request, c.Query("request"))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		overridenCont = &HmacBase64SignedRequest{
			Signature: c.Query(queryStringSignature),
			Request:   []byte(request),
		}
	}
	log.Debug("verifying signature")
	isValid, req, err, code := verifyRequestInternal(c, proxyConfig.HmacSecret, determineIfBase64(c), overridenCont)
	if err != nil {
		c.JSON(code, apiError2.ApiError{
			HttpCode: code,
			Message:  err.Error(),
			Success:  false,
		})
		return
	}
	if !isValid {
		c.JSON(403, apiError2.ApiError{
			HttpCode: 403,
			Message:  stringsInvalidHmacSignature,
			Success:  false,
		})
		return
	}
	performProxyAction(req, c)
}

func determineIfBase64(c *gin.Context) bool {
	log.Debug("determineIfBase64")
	isBase64R := c.Query(queryStringIsBase64)
	isBase64L := c.Query(queryStringIsBase64New)
	isBase64 := false
	if isBase64R == "true" || isBase64L == "true" {
		isBase64 = true
		log.Debug("determineIfBase64 is true")
	}
	return isBase64
}
