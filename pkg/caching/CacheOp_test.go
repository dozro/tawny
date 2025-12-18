package caching

import (
	"errors"
	"testing"
	"time"

	"github.com/dozro/tawny/pkg/musicbrainz_types"
	"github.com/patrickmn/go-cache"
)

const ExampleArtistMBID = "f59c5520-5f46-4d2c-b2c4-822eabf53419"
const ExampleRecMBID = "0815f294-becd-4544-bad0-5d43b7426cce"

func setupTestCache() {
	useLocalInMemoryCaching = true
	musicBrainzRecordingCacheInMemory = cache.New(5*time.Minute, 10*time.Minute)
	musicBrainzArtistCacheInMemory = cache.New(5*time.Minute, 10*time.Minute)
	musicBrainzErrorCacheInMemory = cache.New(5*time.Minute, 10*time.Minute)
}

func TestMusicBrainzCacheRecordingAddAndGet(t *testing.T) {
	setupTestCache()

	mbid := ExampleRecMBID
	recording := musicbrainz_types.Recording{ID: mbid}

	MusicBrainzCacheRecordingAdd(mbid, recording)
	time.Sleep(10 * time.Millisecond) // Wait for goroutine

	data, found, source := MusicBrainzCacheRecordingGet(mbid)
	if !found {
		t.Error("Expected recording to be found in cache")
	}
	if source != "local" {
		t.Errorf("Expected source to be 'local', got '%s'", source)
	}
	if data == nil {
		t.Error("Expected data to not be nil")
	}
}

func TestMusicBrainzCacheRecordingExists(t *testing.T) {
	setupTestCache()

	mbid := ExampleRecMBID
	recording := musicbrainz_types.Recording{ID: mbid}

	if MusicBrainzCacheRecordingExists(mbid) {
		t.Error("Expected recording to not exist initially")
	}

	MusicBrainzCacheRecordingAdd(mbid, recording)
	time.Sleep(10 * time.Millisecond)

	if !MusicBrainzCacheRecordingExists(mbid) {
		t.Error("Expected recording to exist after adding")
	}
}

func TestMusicBrainzCacheRecordingRemove(t *testing.T) {
	setupTestCache()

	mbid := ExampleRecMBID
	recording := musicbrainz_types.Recording{ID: mbid}

	MusicBrainzCacheRecordingAdd(mbid, recording)
	time.Sleep(10 * time.Millisecond)

	MusicBrainzCacheRecordingRemove(mbid)

	if MusicBrainzCacheRecordingExists(mbid) {
		t.Error("Expected recording to not exist after removal")
	}
}

func TestMusicBrainzCacheArtistAddAndGet(t *testing.T) {
	setupTestCache()

	mbid := ExampleArtistMBID
	artist := musicbrainz_types.Artist{ID: mbid}

	MusicBrainzCacheArtistAdd(mbid, artist)
	time.Sleep(10 * time.Millisecond)

	data, found, source := MusicBrainzCacheArtistGet(mbid)
	if !found {
		t.Error("Expected artist to be found in cache")
	}
	if source != "local" {
		t.Errorf("Expected source to be 'local', got '%s'", source)
	}
	if data == nil {
		t.Error("Expected data to not be nil")
	}
}

func TestMusicBrainzCacheArtistExists(t *testing.T) {
	setupTestCache()

	mbid := "artist-mbid-456"
	artist := musicbrainz_types.Artist{ID: mbid}

	if MusicBrainzCacheArtistExists(mbid) {
		t.Error("Expected artist to not exist initially")
	}

	MusicBrainzCacheArtistAdd(mbid, artist)
	time.Sleep(10 * time.Millisecond)

	if !MusicBrainzCacheArtistExists(mbid) {
		t.Error("Expected artist to exist after adding")
	}
}

func TestMusicBrainzCacheArtistRemove(t *testing.T) {
	setupTestCache()

	mbid := ExampleArtistMBID
	artist := musicbrainz_types.Artist{ID: mbid}

	MusicBrainzCacheArtistAdd(mbid, artist)
	time.Sleep(10 * time.Millisecond)

	MusicBrainzCacheArtistRemove(mbid)

	if MusicBrainzCacheArtistExists(mbid) {
		t.Error("Expected artist to not exist after removal")
	}
}

func TestMusicBrainzCacheErrorAddAndGet(t *testing.T) {
	setupTestCache()

	mbid := ExampleArtistMBID
	testError := errors.New("test error")

	MusicBrainzCacheErrorAdd(mbid, testError)
	time.Sleep(10 * time.Millisecond)

	err, found, source := MusicBrainzCacheErrorGet(mbid)
	if !found {
		t.Error("Expected error to be found in cache")
	}
	if source != "local" {
		t.Errorf("Expected source to be 'local', got '%s'", source)
	}
	if err == nil {
		t.Error("Expected error to not be nil")
	}
	if err.Error() != testError.Error() {
		t.Errorf("Expected error '%s', got '%s'", testError.Error(), err.Error())
	}
}

func TestMusicBrainzCacheErrorExists(t *testing.T) {
	setupTestCache()

	mbid := "error-mbid-456"
	testError := errors.New("test error")

	if MusicBrainzCacheErrorExists(mbid) {
		t.Error("Expected error to not exist initially")
	}

	MusicBrainzCacheErrorAdd(mbid, testError)
	time.Sleep(10 * time.Millisecond)

	if !MusicBrainzCacheErrorExists(mbid) {
		t.Error("Expected error to exist after adding")
	}
}

func TestMusicBrainzCacheErrorRemove(t *testing.T) {
	setupTestCache()

	mbid := "error-mbid-789"
	testError := errors.New("test error")

	MusicBrainzCacheErrorAdd(mbid, testError)
	time.Sleep(10 * time.Millisecond)

	MusicBrainzCacheErrorRemove(mbid)

	if MusicBrainzCacheErrorExists(mbid) {
		t.Error("Expected error to not exist after removal")
	}
}

func TestCacheWithCachingDisabled(t *testing.T) {
	useLocalInMemoryCaching = false

	mbid := "disabled-test"
	recording := musicbrainz_types.Recording{ID: mbid}

	MusicBrainzCacheRecordingAdd(mbid, recording)

	if MusicBrainzCacheRecordingExists(mbid) {
		t.Error("Expected recording to not exist when caching is disabled")
	}

	_, found, source := MusicBrainzCacheRecordingGet(mbid)
	if found {
		t.Error("Expected recording to not be found when caching is disabled")
	}
	if source != "none" {
		t.Errorf("Expected source to be 'none', got '%s'", source)
	}
}
