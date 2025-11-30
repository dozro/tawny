package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dozro/tawny/pkg/apiError"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func pageLimitAuthReq(c *gin.Context) (string, string, int, int) {
	apikey := c.Request.Header.Get("Authorization")
	username := c.Param("username")
	var limit, page int
	if c.Query("page") != "" {
		var err error
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			// if atoi fails interpret as -1
			page = -1
		}
	} else if c.Query("start") != "" && c.Query("page") == "-1" {
		var err error
		page, err = strconv.Atoi(c.Query("start"))
		if err != nil {
			// if atoi fails interpret as -1
			page = -1
		}
	} else {
		page = -1
	}
	if c.Query("limit") != "" {
		var err error
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			// if atoi fails interpret as -1
			limit = -1
		}
	} else {
		limit = -1
	}
	return apikey, username, page, limit
}

func fromToAuthReq(c *gin.Context) (string, string, int, int) {
	apikey := c.Request.Header.Get("Authorization")
	username := c.Param("username")
	var to, from int
	if c.Query("from") != "" {
		var err error
		from, err = strconv.Atoi(c.Query("from"))
		if err != nil {
			// if atoi fails interpret as -1
			from = -1
		}
	} else {
		from = -1
	}
	if c.Query("to") != "" {
		var err error
		to, err = strconv.Atoi(c.Query("to"))
		if err != nil {
			// if atoi fails interpret as -1
			to = -1
		}
	} else {
		to = -1
	}
	return apikey, username, from, to
}

func checkIfArrayIsEmpty[T any](c *gin.Context, arr []T, errmsg string) bool {
	if len(arr) == 0 {
		render(c, http.StatusConflict, apiError.ApiError{
			HttpCode:          http.StatusConflict,
			InternalErrorCode: apiError.ArrayIsUnexpectedEmpty,
			InternalErrorMsg:  apiError.ArrayIsUnexpectedEmpty.String(),
			Message:           errmsg,
			Success:           false,
			Date:              time.Now().String(),
		})
		return true
	}
	return false
}

func checkIfAcceptImage(c *gin.Context) bool {
	if !supportedImageTypes.MatchString(c.Request.Header.Get("Accept")) {
		render(c, http.StatusNotAcceptable, apiError.ApiError{
			HttpCode:          http.StatusNotAcceptable,
			InternalErrorCode: apiError.SelectedImageEncodingNotSupported,
			InternalErrorMsg:  apiError.SelectedImageEncodingNotSupported.String(),
			Message:           "The image encoding in your Accept header is not supported",
			Success:           false,
			Date:              time.Now().String(),
		})
		return false
	}
	return true
}

func redirectToHMACEndpoint(c *gin.Context, apiId string, apipara HmacProxyRequestApiParameters) bool {
	signature := c.Query("signature")
	if signature != "" {
		req := HmacProxyRequest{
			Method:        "GET",
			ApiIdentifier: apiId,
			ApiParameters: apipara,
		}
		reqBytes, err := json.Marshal(req)
		if err != nil {
			// this shouldn't happen
			log.Errorf("Unexpected Error in json marshall of HmacProxy Request: %v", err)
			return false
		}
		reqb64 := base64.URLEncoding.EncodeToString(reqBytes)
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/api/v1/hmac/execute?isBase64=true&signature=%s&request=%s", signature, reqb64))
		return true
	}
	return false
}

func render(c *gin.Context, status int, payload interface{}) {
	switch c.GetHeader("Accept") {
	case "application/xml":
		{
			c.XML(status, payload)
			break
		}
	case "application/yaml":
		{
			c.YAML(status, payload)
			break
		}
	default:
		c.JSON(status, payload)
	}
}
