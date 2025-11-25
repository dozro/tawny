package server

import (
	"github.com/dozro/tawny/internal/pkg/proxy_config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var proxyConfig *proxy_config.ProxyConfig

func StartServer(config *proxy_config.ProxyConfig) {

	proxyConfig = config

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(serverHeader)

	api := router.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user")
	user.GET(":username", getUserInfo)
	user.GET(":username/friends", getUserFriends)
	user.GET(":username/tracks/loved", getUserLovedTracks)
	user.GET(":username/tracks/recent", getUserRecentTracks)
	user.GET(":username/tracks/current", getUserCurrentTrack)
	user.GET(":username/tracks/current/embed", getUserCurrentTrackEmbed)
	user.GET(":username/top/albums", getUserTopAlbums)

	router.StaticFile("/swagger.yaml", "./api/apispec.yaml")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger.yaml")))

	hmacapi := v1.Group("/hmac")
	hmacapi.POST("sign", signRequest)
	hmacapi.POST("sign/base64", signBase64Request)
	hmacapi.POST("verify", verifyRequest)
	hmacapi.POST("verify/againstServer", verifyAgainstServerSecret)
	hmacapi.POST("execute", executeSignedRequest)
	hmacapi.GET("execute", executeSignedRequest)

	router.Run()
}
