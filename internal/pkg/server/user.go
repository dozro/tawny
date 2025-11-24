package server

import (
	"github.com/dozro/tawny/internal/pkg/client"

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

func getUserTopAlbums(c *gin.Context) {
	apikey, username, page, limit := pageLimitAuthReq(c)
	if apikeyUndefined(apikey, c) {
		return
	}
	ta, err := client.GetUserTopAlbum(username, apikey, limit, page)
	if handleError(err, c) {
		return
	}
	c.JSON(200, ta)
}

func getUserLovedTracks(c *gin.Context) {
	apikey, username, page, limit := pageLimitAuthReq(c)
	if apikeyUndefined(apikey, c) {
		return
	}
	lt, err := client.GetUserLovedTracks(username, apikey, limit, page)
	if handleError(err, c) {
		return
	}
	c.JSON(200, lt)
}

func getUserRecentTracks(c *gin.Context) {
	apikey, username, page, limit := pageLimitAuthReq(c)
	if apikeyUndefined(apikey, c) {
		return
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
