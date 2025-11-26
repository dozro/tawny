package server

import (
	"github.com/gin-gonic/gin"
	healthcheck "github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	"github.com/tavsec/gin-healthcheck/config"
)

func addHealthChecks(r *gin.Engine) {
	pingCheck := checks.NewPingCheck("https://ws.audioscrobbler.com/2.0/", "HEAD", 10, nil, nil)
	envHmacSecretCheck := checks.NewEnvCheck("TAWNY_HMAC_SECRET")
	envLastFmCheck := checks.NewEnvCheck("TAWNY_LASTFM_API_KEY")
	healthcheck.New(r, config.DefaultConfig(), []checks.Check{pingCheck, envHmacSecretCheck, envLastFmCheck})
}
