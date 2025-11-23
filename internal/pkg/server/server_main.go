package server

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/user/:username", getUserInfo)
	router.GET("/user/:username/tracks/loved", getUserLovedTracks)
	router.GET("/user/:username/tracks/recent", getUserRecentTracks)
	router.Run() // listens on 0.0.0.0:8080 by default
}
