package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func pageLimitAuthReq(c *gin.Context) (string, string, int, int) {
	apikey := c.Request.Header.Get("Authorization")
	username := c.Param("username")
	var limit, page int
	if c.Query("page") != "" {
		page, _ = strconv.Atoi(c.Query("page"))
	} else {
		page = -1
	}
	if c.Query("limit") != "" {
		limit, _ = strconv.Atoi(c.Query("limit"))
	} else {
		limit = -1
	}
	return apikey, username, page, limit
}
