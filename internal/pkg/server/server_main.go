package server

import (
	"fmt"

	"github.com/dozro/tawny/internal/pkg/server_config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var proxyConfig *server_config.ServerConfig

func StartServer(config *server_config.ServerConfig) {

	proxyConfig = config

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(serverHeader)
	if config.ReleaseMode {
		log.Info("Running in ReleaseMode; setting Gin to Release Mode")
		gin.SetMode(gin.ReleaseMode)
	} else if config.DebugMode {
		log.Info("Running in DebugMode; setting Gin to Debug Mode")
		gin.SetMode(gin.DebugMode)
	}

	api := router.Group(config.ApiBasePath)
	v1 := api.Group("/v1")

	user := v1.Group("/user")
	user.Use(disabledEndpointMiddleware())
	un := user.Group(":username")
	untracks := un.Group("tracks")
	unweekly := un.Group("chart/weekly")
	un.GET("", getUserInfo)
	un.GET("friends", getUserFriends)
	untracks.GET("loved", getUserLovedTracks)
	untracks.GET("recent", getUserRecentTracks)
	untracks.GET("current", getUserCurrentTrack)
	untracks.GET("current/embed", getUserCurrentTrackEmbed)

	un.GET("top/albums", getUserTopAlbums)
	un.GET("top/tracks", getUserTopTracks)
	unweekly.GET("album", getUserWeeklyChart)

	router.StaticFile("/swagger.yaml", "./api/apispec.yaml")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger.yaml")))

	hmacapi := v1.Group("/hmac")
	hmacapi.Use(disabledEndpointMiddleware())
	hmacapi.POST("sign", signRequest)
	hmacapi.POST("sign/base64", signBase64Request)
	hmacapi.POST("verify", verifyRequest)
	hmacapi.POST("verify/againstServer", verifyAgainstServerSecret)
	hmacapi.POST("verify/against_server", verifyAgainstServerSecret)
	hmacapi.POST("execute", executeSignedRequest)
	hmacapi.GET("execute", executeSignedRequest)

	musicbrainz := v1.Group("/musicbrainz")
	musicbrainz.Use(disabledEndpointMiddleware())
	musicbrainz.GET("lookup/artist/by_mbid/:artist_mbid", lookUpArtistByMbid)

	addHealthChecks(router)

	router.Run(fmt.Sprintf("%s:%d", proxyConfig.ApiHost, proxyConfig.ApiPort))
}
