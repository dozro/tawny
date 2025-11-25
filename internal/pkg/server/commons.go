package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func pageLimitAuthReq(c *gin.Context) (string, string, int, int) {
	apikey := c.Request.Header.Get("Authorization")
	username := c.Param("username")
	var limit, page int
	if c.Query("page") != "" {
		page, _ = strconv.Atoi(c.Query("page"))
	} else {
		page = -1
	}
	if c.Query("limit") != "" {
		limit, _ = strconv.Atoi(c.Query("limit"))
	} else {
		limit = -1
	}
	return apikey, username, page, limit
}

func redirectToHMACEndpoint(c *gin.Context, apiId string, apipara HmacProxyRequestApiParameters) bool {
	signature := c.Query("signature")
	if signature != "" {
		req := HmacProxyRequest{
			Method:        "GET",
			ApiIdentifier: apiId,
			ApiParameters: apipara,
		}
		reqBytes, _ := json.Marshal(req)
		reqb64 := base64.URLEncoding.EncodeToString(reqBytes)
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/api/v1/hmac/execute?isBase64=true&signature=%s&request=%s", signature, reqb64))
		return true
	}
	return false
}
