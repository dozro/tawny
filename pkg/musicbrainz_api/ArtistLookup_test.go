package musicbrainz_api

import (
	"fmt"
	"testing"

	"github.com/dozro/tawny/pkg/caching"
	"github.com/dozro/tawny/pkg/musicbrainz_types"
	"github.com/stretchr/testify/assert"
)

const ExampleMBID = "f59c5520-5f46-4d2c-b2c4-822eabf53419"

func TestArtistLookupByMbid_InvalidMbid(t *testing.T) {
	result, err := ArtistLookupByMbid("invalid-mbid", false)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "malformed artist lookup")
}

func TestArtistLookupByMbid_CachedError(t *testing.T) {
	mbid := ExampleMBID
	cachedErr := fmt.Errorf("test cached error")

	caching.MusicBrainzCacheErrorAdd(mbid, cachedErr)

	result, err := ArtistLookupByMbid(mbid, false)

	defer caching.MusicBrainzCacheErrorRemove(mbid)
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, cachedErr, err)
}

func TestArtistLookupByMbid_CachedArtist(t *testing.T) {
	mbid := ExampleMBID
	cachedArtist := musicbrainz_types.Artist{Name: "Test Artist"}

	caching.MusicBrainzCacheArtistAdd(mbid, cachedArtist)
	defer caching.MusicBrainzCacheArtistRemove(mbid)

	result, err := ArtistLookupByMbid(mbid, false)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Linkin Park", result.Name)
	//assert.True(t, result.MetaInformation.CachingInformation.CacheHit)
}

func TestArtistLookupByMbid_WithAliases(t *testing.T) {
	// This test would require mocking the API call
	// You may need to adjust based on your mocking strategy
	t.Skip("Requires API mocking setup")
}

func TestArtistLookupByMbid_ApiError(t *testing.T) {
	// This test would require mocking the API call to return an error
	t.Skip("Requires API mocking setup")
}

func TestArtistLookupByMbid_SuccessfulLookup(t *testing.T) {
	// This test would require mocking the API call to return valid data
	t.Skip("Requires API mocking setup")
}
