package server

import (
	"fmt"
	"net/http"

	"github.com/dozro/tawny/internal/pkg/apiError"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func disabledEndpointHandler(c *gin.Context) {
	log.Infof("Request made to server disabled endpoint")
	render(c, 503, apiError.ApiError{
		HttpCode:          http.StatusServiceUnavailable,
		Message:           fmt.Sprintf("The endpoint %s is currently disabled on %s by server configuration", c.Request.RequestURI, c.Request.Host),
		InternalErrorCode: apiError.EndpointDisabledByConfig,
		InternalErrorMsg:  "This endpoint is currently disabled by server configuration",
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
