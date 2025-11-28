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

}
