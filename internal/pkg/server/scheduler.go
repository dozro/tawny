package server

import (
	"time"

	"github.com/dozro/tawny/pkg/caching"
	"github.com/roylee0704/gron"
	log "github.com/sirupsen/logrus"
)

var schedule *gron.Cron

func setupScheduler() {
	schedule = gron.New()
	if proxyConfig.CachingConfig.SaveCacheToFS {
		schedule.AddFunc(gron.Every(5*time.Minute), func() {
			log.Info("[Scheduled Task]: Starting save process to fs")
			go caching.SaveToFS(proxyConfig.CachingConfig.LocalCachePath)
		})
	}
	schedule.Start()
}
