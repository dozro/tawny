package main

import (
	"github.com/dozro/tawny/internal/pkg/server"
	"github.com/dozro/tawny/internal/pkg/server_config"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting fm-proxy")
	config := server_config.SetupServerConfig()
	if config.DebugMode {
		log.SetLevel(log.DebugLevel)
	} else if config.ReleaseMode {
		log.SetLevel(log.InfoLevel)
	}
	server.StartServer(config)
}
