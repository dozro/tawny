package server

import (
	"lastfm-proxy/internal/pkg/client"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func getUserInfo(c *gin.Context) {
	apikey := c.Request.Header.Get("Authorization")
	username := c.Param("username")
	log.Infof("getUserInfo: %s, %s", username, apikey)
	if apikeyUndefined(apikey, c) {
		return
	}
	userinfo, err := client.GetUserInfo(username, apikey)
	if handleError(err, c) {
		return
	}
	c.JSON(200, userinfo)
}

func getUserLovedTracks(c *gin.Context) {
	apikey := c.Request.Header.Get("Authorization")
	username := c.Param("username")
	if apikeyUndefined(apikey, c) {
		return
	}
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
	lt, err := client.GetUserLovedTracks(username, apikey, limit, page)
	if handleError(err, c) {
		return
	}
	c.JSON(200, lt)
}

func getUserRecentTracks(c *gin.Context) {
	apikey := c.Request.Header.Get("Authorization")
	username := c.Param("username")
	if apikeyUndefined(apikey, c) {
		return
	}
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
	lt, err := client.GetUserRecentTracks(username, apikey, limit, page)
	if handleError(err, c) {
		return
	}
	c.JSON(200, lt)
}

func getUserCurrentTrack(c *gin.Context) {
	apikey := c.Request.Header.Get("Authorization")
	username := c.Param("username")
	if apikeyUndefined(apikey, c) {
		return
	}
	ct, err := client.GetUserCurrentTrack(username, apikey)
	if handleError(err, c) {
		return
	}
	c.JSON(200, ct)
}
