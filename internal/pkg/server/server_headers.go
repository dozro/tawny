package server

import "github.com/gin-gonic/gin"

func serverHeader(c *gin.Context) {
	c.Header("Server", "Tawny")
	c.Header("X-DNS-Prefetch-Control", "off")
	c.Header("X-Powered-By", "Tawny")
	c.Header("X-XSS-Protection", "1")
	c.Header("X-Frame-Options", "ALLOW")
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
}
