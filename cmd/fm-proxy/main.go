package main

import (
	"lastfm-proxy/internal/pkg/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting fm-proxy")
	log.SetLevel(log.DebugLevel)
	server.StartServer()
}
