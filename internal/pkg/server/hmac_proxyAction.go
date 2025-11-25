package server

import (
	"net/http"

	"github.com/dozro/tawny/internal/pkg/apiError"
	"github.com/dozro/tawny/internal/pkg/client"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func performProxyAction(request *HmacProxyRequest, c *gin.Context) {
	log.Debug("performing proxy action")
	switch {
	case userRegex.MatchString(request.ApiIdentifier):
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
}
