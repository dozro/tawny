package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func serveSwagger(router *gin.Engine) {
	if !proxyConfig.ExtendedServerConfig.RunningInDocker {
		router.StaticFS("/openapi/", http.Dir("./api/openapi"))
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/openapi/openapi.yaml")))
	} else {
		router.StaticFile("/swagger.yaml", "./api/openapi-bundled.yaml")
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger.yaml")))
	}
}
