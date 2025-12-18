package musicbrainz_api

import (
	"fmt"
	"testing"

	"github.com/dozro/tawny/pkg/caching"
	"github.com/dozro/tawny/pkg/musicbrainz_types"
	"github.com/stretchr/testify/assert"
)

const exampleMBID = "f59c5520-5f46-4d2c-b2c4-822eabf53419"
const exampleArtistName = "Linkin Park"
const malformedArtistErrMsg = "malformed artist lookup"
const skipMsg = "Requires API mocking setup"

func TestArtistLookupByMbidInvalidMbid(t *testing.T) {
	result, err := ArtistLookupByMbid("invalid-mbid", false)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), malformedArtistErrMsg)
}

func TestArtistLookupByMbidCachedError(t *testing.T) {
	mbid := exampleMBID
	cachedErr := fmt.Errorf("test cached error")

	caching.MusicBrainzCacheErrorAdd(mbid, cachedErr)

	result, err := ArtistLookupByMbid(mbid, false)

	defer caching.MusicBrainzCacheErrorRemove(mbid)
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, cachedErr, err)
}

func TestArtistLookupByMbidCachedArtist(t *testing.T) {
	mbid := exampleMBID
	cachedArtist := musicbrainz_types.Artist{Name: exampleArtistName, ID: exampleMBID}

	caching.MusicBrainzCacheArtistAdd(mbid, cachedArtist)
	defer caching.MusicBrainzCacheArtistRemove(mbid)

	result, err := ArtistLookupByMbid(mbid, false)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, exampleArtistName, result.Name)
	//assert.True(t, result.MetaInformation.CachingInformation.CacheHit)
}

func TestArtistLookupByMbidWithAliases(t *testing.T) {
	// This test would require mocking the API call
	// You may need to adjust based on your mocking strategy
	t.Skip(skipMsg)
}

func TestArtistLookupByMbidApiError(t *testing.T) {
	// This test would require mocking the API call to return an error
	t.Skip(skipMsg)
}

func TestArtistLookupByMbidSuccessfulLookup(t *testing.T) {
	// This test would require mocking the API call to return valid data
	t.Skip(skipMsg)
}
