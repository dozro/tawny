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
	un := parseUsernameWithoutRet(username)
	username = un.username
	return apikey, username, page, limit
}

func fromToAuthReq(c *gin.Context) (string, string, int, int) {
	apikey := c.Request.Header.Get("Authorization")
	username := c.Param("username")
	un := parseUsernameWithoutRet(username)
	username = un.username
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

func isUsernameCannonical(username string) bool {
	return canonicalUsernameRegexp.MatchString(username)
}

func isLastfmUser(username string) bool {
	if isUsernameCannonical(username) {
		a := canonicalUsernameRegexp.FindStringSubmatch(username)
		log.Debugf("checking if %s contains last.fm (%s)", username, a[1])
		return a[3] == "last.fm"
	}
	return true
}

func isListenBrainzUser(username string) bool {
	if isUsernameCannonical(username) {
		a := canonicalUsernameRegexp.FindStringSubmatch(username)
		log.Debugf("checking if %s contains listenbrainz.org (%s)", username, a[1])
		return a[3] == "listenbrainz.org"
	}
	return false
}

func extractUsernameFromCannonical(username string) string {
	a := canonicalUsernameRegexp.FindStringSubmatch(username)
	if len(a) < 2 {
		log.Debugf("[compat] Unable to extract username from %s; is it a cannonical username?", username)
		return username
	}
	log.Debugf("extracting %s from %s", a[2], username)
	return a[2]
}

type canUser struct {
	Lfm      bool
	Lb       bool
	username string
}

func parseUsernameWithoutRet(username string) canUser {
	isLb := false
	isLfm := false
	un := extractUsernameFromCannonical(username)
	if !isUsernameCannonical(username) {
		log.Warnf("[compat] username %s is not a cannonical username, interpreting as last.fm username", un)
		un = username
		isLfm = true
	} else if isListenBrainzUser(username) {
		log.Debugf("username %s is listenbrainz.org User", un)
		isLb = true
	} else if isLastfmUser(username) {
		log.Debugf("username %s is last.fm User", un)
		isLfm = true
	} else {
		log.Warnf("[compat] username %s was detected as canonical but doesn't match any supported endpoints, interpreting as last.fm username", un)
		un = username
		isLfm = true
	}
	return canUser{
		Lfm:      isLfm,
		Lb:       isLb,
		username: un,
	}
}

func parseUsername(c *gin.Context, username string) (string, bool) {
	if !isLastfmUser(username) {
		render(c, http.StatusBadRequest, apiError.ApiError{
			HttpCode:          http.StatusBadRequest,
			InternalErrorCode: 0,
			InternalErrorMsg:  "",
			Message:           fmt.Sprintf("%s isn't lastfm user", username),
			Data:              nil,
			Success:           false,
			Date:              time.Now().String(),
		})
		return "", false
	}
	if isUsernameCannonical(username) {
		return extractUsernameFromCannonical(username), true
	}
	return "", false
}
