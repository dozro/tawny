package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServer() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/api/v1/user/:username", getUserInfo)
	router.GET("/api/v1/user/:username/tracks/loved", getUserLovedTracks)
	router.GET("/api/v1/user/:username/tracks/recent", getUserRecentTracks)
	router.GET("/api/v1/user/:username/tracks/current", getUserCurrentTrack)

	router.StaticFile("/swagger.yaml", "./api/apispec.yaml")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger.yaml")))

	router.GET("/api/v1/hmac/sign", signRequest)
	router.HEAD("/api/v1/hmac/verify", verifyRequest)
	router.GET("/api/v1/hmac/verify", verifyRequest)
	router.GET("/api/v1/hmac/execute", signRequest)
	router.Run() // listens on 0.0.0.0:8080 by default
}
