package server

import (
	"fmt"
	"net/http"
	"time"

	apiError2 "github.com/dozro/tawny/pkg/apiError"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func disabledEndpointHandler(c *gin.Context) {
	log.Infof("Request made to server disabled endpoint")
	render(c, 503, apiError2.ApiError{
		HttpCode:          http.StatusServiceUnavailable,
		Message:           fmt.Sprintf("The endpoint %s is currently disabled on %s by server configuration", c.Request.RequestURI, c.Request.Host),
		InternalErrorCode: apiError2.EndpointDisabledByConfig,
		InternalErrorMsg:  "This endpoint is currently disabled by server configuration",
		Data:              c.Request,
		Success:           false,
		Date:              time.Now().String(),
	})
	c.Abort()
}

func checkDisabledEndpoint(path string, c *gin.Context) bool {
	// Enable Only HMAC
	if proxyConfig.DisabledEndpoints.EnableOnlyHMACEndpoints &&
		!middlewareHMACEndpointRegex.MatchString(path) {
		disabledEndpointHandler(c)
		return true
	}

	// Disable Image Embedded
	if proxyConfig.DisabledEndpoints.DisableImageEmbeddedEndpoints &&
		middlewareEmbedEndpointRegex.MatchString(path) {
		disabledEndpointHandler(c)
		return true
	}

	// Disable HMAC Signing
	if proxyConfig.DisabledEndpoints.DisableHMACSigningEndpoint &&
		middlewareHMACSignEndpointRegex.MatchString(path) {
		disabledEndpointHandler(c)
		return true
	}

	// Disable MusicBrainz
	if proxyConfig.DisabledEndpoints.DisableMusicBrainzEndpoints &&
		middlewareMusicBrainzEndpointRegex.MatchString(path) {
		disabledEndpointHandler(c)
		return true
	}

	// Nichts blockiert
	return false
}

func disabledEndpointMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if checkDisabledEndpoint(path, c) {
			return
		}
	}
}
