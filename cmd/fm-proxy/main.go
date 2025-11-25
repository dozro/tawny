package main

import (
	"github.com/dozro/tawny/internal/pkg/proxy_config"
	"github.com/dozro/tawny/internal/pkg/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting fm-proxy")
	log.SetLevel(log.DebugLevel)
	config := proxy_config.GetProxyConfig()
	server.StartServer(config)
}
