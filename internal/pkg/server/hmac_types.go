package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type HmacSignedRequest struct {
	Signature string `json:"signature"`
	// Raw JSON to avoid issues with encoding while verifying the request
	Request json.RawMessage `json:"request"`
}

type HmacBase64SignedRequest struct {
	Signature string `json:"signature"`
	Request   []byte `json:"request"`
}

func HmacSignedRequestToBase64(request HmacSignedRequest) HmacBase64SignedRequest {
	signedRequest := HmacBase64SignedRequest{
		Signature: request.Signature,
		Request:   request.Request,
	}
	return signedRequest
}

func Base64ToHmacSignedRequest(b64 HmacBase64SignedRequest) (HmacSignedRequest, error) {
	log.Debugf("decoding base64 signed request: %s", b64)

	var decoded []byte
	if json.Valid(b64.Request) {
		log.Debugf("data is already decoded, will skip base64 decoding")
		decoded = b64.Request
	} else {
		// Base64 decode the JSON request payload
		var err error
		decoded, err = base64.StdEncoding.DecodeString(string(b64.Request))
		if err != nil {
			e := fmt.Errorf("failed to decode base64 request: %w", err)
			log.Error(e)
			return HmacSignedRequest{}, e
		}
	}

	if !json.Valid(decoded) {
		e := fmt.Errorf("decoded request is not valid JSON: %s", decoded)
		log.Error(e)
		return HmacSignedRequest{}, e
	}

	signedReq := HmacSignedRequest{
		Signature: b64.Signature,
		Request:   json.RawMessage(decoded),
	}

	log.Debugf("decoded base64 request: %s, %s", signedReq.Signature, signedReq.Request)
	return signedReq, nil
}

type HmacProxyRequest struct {
	Method        string                        `json:"method"`
	ApiIdentifier string                        `json:"api_identifier"`
	ApiParameters HmacProxyRequestApiParameters `json:"api_parameters"`
	Body          string                        `json:"body"`
}

type HmacProxyRequestApiParameters struct {
	Username string `json:"username"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
}

type HmacProxyResponseWrapper struct {
	HttpCode int               `json:"http_code"`
	Response HmacProxyResponse `json:"response"`
}

type HmacProxyResponse struct {
	Schema string          `json:"$schema"`
	Data   json.RawMessage `json:"data"`
}
