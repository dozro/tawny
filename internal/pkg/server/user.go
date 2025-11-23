package server

import (
	"lastfm-proxy/internal/pkg/client"

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
	c.JSON(200, gin.H{
		"code": 200,
		"data": userinfo,
	})
}
