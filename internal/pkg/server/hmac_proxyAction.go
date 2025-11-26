package server

import (
	"fmt"
	"net/http"

	"github.com/dozro/tawny/internal/pkg/apiError"
	"github.com/dozro/tawny/internal/pkg/client"
	"github.com/dozro/tawny/internal/pkg/embed"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func performProxyAction(request *HmacProxyRequest, c *gin.Context) {
	log.Debug("performing proxy action")
	switch {
	case userInfoRegex.MatchString(request.ApiIdentifier):
		{
			log.Debug("proxy action: user.GetInfo")
			userinfo, err := client.GetUserInfo(request.ApiParameters.Username, proxyConfig.LastFMAPIKey)
			if handleError(err, c) {
				return
			}
			if userinfo == nil {
				c.JSON(http.StatusServiceUnavailable, apiError.ApiError{
					HttpCode:          http.StatusServiceUnavailable,
					InternalErrorCode: 0,
					InternalErrorMsg:  "",
					Message:           "userinfo is nil",
					Data:              request,
					Success:           false,
				})
				return
			}
			log.Debugf("proxy action result: user.GetInfo: %v", userinfo)
			c.JSON(200, userinfo)
			return
		}
	case userNowPlayingRegex.MatchString(request.ApiIdentifier):
		{
			log.Debug("proxy action: user.NowPlaying")
			usernowplaying, err := client.GetUserCurrentTrack(request.ApiParameters.Username, proxyConfig.LastFMAPIKey)
			if handleError(err, c) {
				return
			}
			c.JSON(200, usernowplaying)
			return
		}
	case userNowPlayingEmbed.MatchString(request.ApiIdentifier):
		{
			log.Debug("proxy action: user.NowPlayingEmbed")
			ct, err := client.GetUserCurrentTrack(request.ApiParameters.Username, proxyConfig.LastFMAPIKey)
			if ct == nil || err != nil {
				e := fmt.Errorf("Unexpected or error", err)
				log.Error(e)
				c.AbortWithError(http.StatusInternalServerError, e)
				return
			}
			img, err := embed.EmbedNowPlaying(ct.Track[0].Name, ct.Track[0].Artist.Name, ct.Track[0].Album, ct.Track[0].Image, request.ApiParameters.Username, ct.Track[0].NowPlaying)
			if handleError(err, c) {
				return
			}
			c.Data(http.StatusOK, "image/png", img.Bytes())
		}
	case userRecentlyPlayedRegex.MatchString(request.ApiIdentifier):
		{
			log.Debug("proxy action: user.RecentlyPlayed")
			rp, err := client.GetUserRecentTracks(request.ApiParameters.Username, proxyConfig.LastFMAPIKey, request.ApiParameters.Limit, request.ApiParameters.Page)
			if handleError(err, c) {
				return
			}
			c.JSON(200, rp)
			return
		}
	}
}
