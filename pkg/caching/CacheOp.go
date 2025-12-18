package caching

import (
	"github.com/dozro/tawny/pkg/musicbrainz_types"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

func MusicBrainzCacheRecordingAdd(mbid string, recording musicbrainz_types.Recording) {
	if useLocalInMemoryCaching {
		log.Debugf("adding item with mbid=%s to cache", mbid)
		go musicBrainzRecordingCacheInMemory.Set(mbid, recording, cache.DefaultExpiration)
	}
}

func MusicBrainzCacheErrorAdd(mbid string, err error) {
	log.Debugf("logging error for mbid=%s with error=%v", mbid, err)
	if useLocalInMemoryCaching {
		log.Debugf("logging error for mbid=%s to cache", mbid)
		go musicBrainzErrorCacheInMemory.Set(mbid, err, cache.DefaultExpiration)
	}
}

func MusicBrainzCacheErrorExists(mbid string) bool {
	if useLocalInMemoryCaching {
		dat, e := musicBrainzErrorCacheInMemory.Get(mbid)
		return dat != nil || e
	}
	return false
}

func MusicBrainzCacheErrorGet(mbid string) (error, bool, string) {
	if useLocalInMemoryCaching {
		d, e := musicBrainzRecordingCacheInMemory.Get(mbid)
		if d == nil {
			return nil, false, "local"
		}
		return d.(error), e, "local"
	}
	return nil, false, "none"
}

func MusicBrainzCacheRecordingExists(mbid string) bool {
	if useLocalInMemoryCaching {
		dat, e := musicBrainzRecordingCacheInMemory.Get(mbid)
		return dat != nil || e
	}
	return false
}

func MusicBrainzCacheArtistExists(mbid string) bool {
	if useLocalInMemoryCaching {
		dat, e := musicBrainzArtistCacheInMemory.Get(mbid)
		return dat != nil || e
	}
	return false
}

func MusicBrainzCacheRecordingGet(mbid string) (interface{}, bool, string) {
	if useLocalInMemoryCaching {
		d, e := musicBrainzRecordingCacheInMemory.Get(mbid)
		return d, e, "local"
	}
	return nil, false, "none"
}

func MusicBrainzCacheArtistGet(mbid string) (interface{}, bool, string) {
	if useLocalInMemoryCaching {
		d, e := musicBrainzArtistCacheInMemory.Get(mbid)
		return d, e, "local"
	}
	return nil, false, "none"
}

func MusicBrainzCacheArtistAdd(mbid string, artist musicbrainz_types.Artist) {
	if useLocalInMemoryCaching {
		go musicBrainzArtistCacheInMemory.Set(mbid, artist, cache.DefaultExpiration)
	}
}
