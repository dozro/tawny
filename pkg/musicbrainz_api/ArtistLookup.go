package musicbrainz_api

import (
	"fmt"

	"github.com/dozro/tawny/internal/pkg/api_commons"
	"github.com/dozro/tawny/pkg/caching"
	"github.com/dozro/tawny/pkg/common_types"
	"github.com/dozro/tawny/pkg/musicbrainz_types"
	log "github.com/sirupsen/logrus"
)

func ArtistLookupByMbid(mbid string, includeAliases bool) (*musicbrainz_types.Artist, error) {
	if !mbidRegexp.MatchString(mbid) {
		return nil, fmt.Errorf("musicbrainz_api: malformed artist lookup for mbid=%s", mbid)
	}
	if caching.MusicBrainzCacheErrorExists(mbid) {
		log.Warnf("found cached error for %s. Skipping lookup.", mbid)
		d, _, _ := caching.MusicBrainzCacheErrorGet(mbid)
		if d == nil {
			d = fmt.Errorf("cached error is faulty. This most likely is a bug.")
		}
		return nil, d
	}

	if caching.MusicBrainzCacheArtistExists(mbid) {
		log.Infof("Artist with mbid=%s exists in cache; getting it and skipping further lookup", mbid)
		d, _, b := caching.MusicBrainzCacheArtistGet(mbid)
		dm := d.(musicbrainz_types.Artist)
		dm.MetaInformation = common_types.CreateMetaInformation(true, b)
		return &dm, nil
	}

	apiUrl := fmt.Sprintf("https://musicbrainz.org/ws/2/artist/%s", mbid)
	if includeAliases {
		apiUrl += "?inc=aliases"
	}
	data, err := api_commons.FetchXML[musicbrainz_types.WrappedArtistLookupResult](apiUrl)
	if err != nil {
		caching.MusicBrainzCacheErrorAdd(mbid, err)
		return nil, err
	}
	data.Artist.MetaInformation = common_types.CreateMetaInformation(false, false)
	caching.MusicBrainzCacheArtistAdd(mbid, data.Artist)
	return &data.Artist, nil
}
