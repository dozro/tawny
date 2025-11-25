package server

import (
	"github.com/dozro/tawny/internal/pkg/client"
	"github.com/gin-gonic/gin"
)

func performProxyAction(request *HmacProxyRequest, c *gin.Context) {
	switch request.ApiIdentifier {
	case "/user":
		userinfo, err := client.GetUserInfo(request.ApiParameters.Username, proxyConfig.LastFMAPIKey)
		if handleError(err, c) {
			return
		}
		c.JSON(200, userinfo)
	}
}
