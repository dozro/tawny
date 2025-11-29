package server

import (
	"fmt"
	"net/http"

	"github.com/dozro/tawny/internal/pkg/client"
	"github.com/dozro/tawny/internal/pkg/embed"
	"github.com/dozro/tawny/internal/pkg/security"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func getUserInfo(c *gin.Context) {
	apikey := c.Request.Header.Get("Authorization")
	username := c.Param("username")
	log.Infof("getUserInfo: %s, %s", username, security.MaskAPIKey(apikey))
	if apikeyUndefined(apikey, c) {
		return
	}
	userinfo, err := client.LfmUserInfo(username, apikey)
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
	ta, err := client.LfmUserTopAlbum(username, apikey, limit, page)
	if handleError(err, c) {
		return
	}
	c.JSON(200, ta)
}

func getUserLovedTracks(c *gin.Context) {
	apikey, username, page, limit := pageLimitAuthReq(c)
	if redirectToHMACEndpoint(c, "/user/tracks/loved", HmacProxyRequestApiParameters{Username: username}) {
		return
	}
	if apikeyUndefined(apikey, c) {
		return
	}
	lt, err := client.UserLovedTracks(username, apikey, limit, page)
	if handleError(err, c) {
		return
	}
	c.JSON(200, lt)
}

func getUserRecentTracks(c *gin.Context) {
	apikey, username, page, limit := pageLimitAuthReq(c)
	embedMusicBrainz := c.Query("fetch_musicbrainz")
	embedMusicBrainzB := false
	if embedMusicBrainz == "true" {
		embedMusicBrainzB = true
	}
	if redirectToHMACEndpoint(c, "/user/tracks/recent", HmacProxyRequestApiParameters{Username: username}) {
		return
	}
	if apikeyUndefined(apikey, c) {
		return
	}
	lt, err := client.UserRecentTracks(username, apikey, limit, page, embedMusicBrainzB)
	if handleError(err, c) {
		return
	}
	c.JSON(200, lt)
}

func getUserCurrentTrack(c *gin.Context) {
	apikey := c.Request.Header.Get("Authorization")
	username := c.Param("username")
	if redirectToHMACEndpoint(c, "/user/tracks/current", HmacProxyRequestApiParameters{Username: username}) {
		return
	}
	if apikeyUndefined(apikey, c) {
		return
	}
	ct, err := client.LfmUserCurrentTrack(username, apikey)
	if handleError(err, c) {
		return
	}
	c.JSON(200, ct)
}

func getUserCurrentTrackEmbed(c *gin.Context) {
	apikey := c.Request.Header.Get("Authorization")
	username := c.Param("username")
	if apikeyUndefined(apikey, c) {
		return
	}
	ct, err := client.LfmUserCurrentTrack(username, apikey)
	if handleError(err, c) {
		return
	}
	img, err := embed.EmbedNowPlaying(ct.Track[0].Name, ct.Track[0].Artist.Name, ct.Track[0].Album, ct.Track[0].Image, username, ct.Track[0].NowPlaying)
	if handleError(err, c) {
		return
	}
	if img == nil {
		e := fmt.Errorf("image not found")
		log.Error("no valid image data found")
		c.AbortWithError(http.StatusInternalServerError, e)
		return
	}
	c.Data(http.StatusOK, "image/png", img.Bytes())
}

func getUserFriends(c *gin.Context) {
	apikey, username, page, limit := pageLimitAuthReq(c)
	if apikeyUndefined(apikey, c) {
		return
	}
	recentTracks := c.Query("get_recent_tracks")
	var getRT = false
	if recentTracks == "true" {
		getRT = true
	}
	if apikeyUndefined(apikey, c) {
		return
	}
	uf, err := client.LfmUserFriends(username, apikey, limit, page, getRT)
	if handleError(err, c) {
		return
	}
	render(c, http.StatusOK, uf)
}

func getUserTopTracks(c *gin.Context) {
	apikey, username, page, limit := pageLimitAuthReq(c)
	if apikeyUndefined(apikey, c) {
		return
	}
	tt, err := client.LfmUserTopTracks(username, apikey, limit, page)
	if handleError(err, c) {
		return
	}
	render(c, http.StatusOK, tt)
}
