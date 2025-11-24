package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServer() {
	router := gin.Default()
	router.Use(gin.Recovery())

	api := router.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user")
	user.GET(":username", getUserInfo)
	user.GET(":username/friends", getUserFriends)
	user.GET(":username/tracks/loved", getUserLovedTracks)
	user.GET(":username/tracks/recent", getUserRecentTracks)
	user.GET(":username/tracks/current", getUserCurrentTrack)
	user.GET(":username/top/albums", getUserTopAlbums)

	router.StaticFile("/swagger.yaml", "./api/apispec.yaml")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger.yaml")))

	hmacapi := v1.Group("/hmac")
	hmacapi.GET("sign", signRequest)
	hmacapi.HEAD("verify", verifyRequest)
	hmacapi.GET("verify", verifyRequest)
	hmacapi.GET("execute", signRequest)
	router.Run() // listens on 0.0.0.0:8080 by default
}
