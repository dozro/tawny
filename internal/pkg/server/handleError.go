package server

import (
	"time"

	apiError2 "github.com/dozro/tawny/pkg/apiError"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func handleError(err error, c *gin.Context) bool {
	if err != nil {
		apiErr := apiError2.ApiError{
			Data: gin.H{
				"request": gin.H{
					"host":   c.Request.Host,
					"path":   c.Request.URL.Path,
					"query":  c.Request.URL.RawQuery,
					"method": c.Request.Method,
				},
			},
			Message:           err.Error(),
			HttpCode:          401,
			Success:           false,
			Date:              time.Now().String(),
			InternalErrorMsg:  apiError2.InternalTawnyError.String(),
			InternalErrorCode: apiError2.InternalTawnyError,
		}
		render(c, 500, apiErr)
		log.Warnf("A server error was caught during the request process: %s", err.Error())
		return true
	}
	return false
}
