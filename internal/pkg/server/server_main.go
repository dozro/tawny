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
	user.GET(":username", getUserInfo)
	user.GET(":username/friends", getUserFriends)
	user.GET(":username/tracks/loved", getUserLovedTracks)
	user.GET(":username/tracks/recent", getUserRecentTracks)
	user.GET(":username/tracks/current", getUserCurrentTrack)
	user.GET(":username/tracks/current/embed", getUserCurrentTrackEmbed)
	user.GET(":username/top/albums", getUserTopAlbums)
	user.GET(":username/top/tracks", getUserTopTracks)

	router.StaticFile("/swagger.yaml", "./api/apispec.yaml")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger.yaml")))

	hmacapi := v1.Group("/hmac")
	hmacapi.POST("sign", signRequest)
	hmacapi.POST("sign/base64", signBase64Request)
	hmacapi.POST("verify", verifyRequest)
	hmacapi.POST("verify/againstServer", verifyAgainstServerSecret)
	hmacapi.POST("verify/against_server", verifyAgainstServerSecret)
	hmacapi.POST("execute", executeSignedRequest)
	hmacapi.GET("execute", executeSignedRequest)

	musicbrainz := v1.Group("/musicbrainz")
	musicbrainz.GET("lookup/artist/by_mbid/:artist_mbid", lookUpArtistByMbid)

	addHealthChecks(router)

	router.Run(fmt.Sprintf("%s:%d", proxyConfig.ApiHost, proxyConfig.ApiPort))
}
