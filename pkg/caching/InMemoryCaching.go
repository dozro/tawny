package caching

import (
	"fmt"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

var musicBrainzRecordingCacheInMemory *cache.Cache
var musicBrainzArtistCacheInMemory *cache.Cache
var musicBrainzErrorCacheInMemory *cache.Cache

func SetupLocalCache(basePath string) {
	log.Debug("setting up local cache")
	useLocalInMemoryCaching = true
	setupForMusicBrainz()
	go LoadFromFS(basePath)
}

func setupForMusicBrainz() {
	log.Debug("setting up local cache for MusicBrainz")
	musicBrainzRecordingCacheInMemory = cache.New(30*time.Minute, 3*time.Hour)
	musicBrainzArtistCacheInMemory = cache.New(30*time.Minute, 24*time.Hour)
	musicBrainzErrorCacheInMemory = cache.New(5*time.Minute, 15*time.Minute)
}

func LoadFromFS(basePath string) {
	log.Info("loading cache from fs")
	if musicBrainzArtistCacheInMemory.LoadFile(fmt.Sprintf("%s/musicBrainzArtist.cache", basePath)) != nil || musicBrainzRecordingCacheInMemory.LoadFile(fmt.Sprintf("%s/musicBrainzArtist.cache", basePath)) != nil {
		log.Warn("error loading cache from fs")
	}

}

func SaveToFS(basePath string) {
	log.Info("persisting cache to filesystem")
	err := os.MkdirAll(basePath, 0755)
	if err != nil {
		log.Error(err)
		return
	}
	if musicBrainzArtistCacheInMemory.SaveFile(fmt.Sprintf("%s/musicBrainzArtist.cache", basePath)) != nil {
		log.Error("error while saving mb artists cache")
		return
	}
	if musicBrainzRecordingCacheInMemory.SaveFile(fmt.Sprintf("%s/musicBrainzRecording.cache", basePath)) != nil {
		log.Error("error while saving mb recordings cache")
		return
	}
}
