package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func handleError(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(500, gin.H{
			"code":  500,
			"error": err,
			"data": gin.H{
				"request": gin.H{
					"host":   c.Request.Host,
					"path":   c.Request.URL.Path,
					"query":  c.Request.URL.RawQuery,
					"method": c.Request.Method,
				},
			},
		})
		log.Warnf("A server error was caught during the request process: %s", err.Error())
	}
	return false
}
