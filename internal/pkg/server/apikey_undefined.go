package server

import "github.com/gin-gonic/gin"

func apikeyUndefined(apikey string, c *gin.Context) bool {
	if apikey == "" {
		c.JSON(401, gin.H{
			"code":  401,
			"error": "apikey undefined",
		})
		return true
	}
	return false
}
