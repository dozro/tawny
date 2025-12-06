package main

import (
	"github.com/dozro/tawny/internal/pkg/server"
	"github.com/dozro/tawny/internal/pkg/server_config"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting fm-proxy")
	config := server_config.SetupServerConfig()
	if config.ExtendedServerConfig.RunningInDocker {
		log.Infof("This server is running on tawny version %s with revision %s", config.ExtendedServerConfig.TawnyVersion, config.ExtendedServerConfig.TawnyRevision)
	}
	if config.DebugMode {
		log.SetLevel(log.DebugLevel)
	} else if config.ReleaseMode {
		log.SetLevel(log.InfoLevel)
	}
	if config.ExtendedServerConfig.LogOutputFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	server.StartServer(config)
}
