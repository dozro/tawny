package musicbrainz_api

import (
	"fmt"

	"github.com/dozro/tawny/internal/pkg/api_commons"
	"github.com/dozro/tawny/pkg/caching"
	"github.com/dozro/tawny/pkg/common_types"
	"github.com/dozro/tawny/pkg/musicbrainz_types"
	log "github.com/sirupsen/logrus"
)

func RecordingLookupByMbid(mbid string, includeAliases bool) (*musicbrainz_types.Recording, error) {
	if !mbidRegexp.MatchString(mbid) {
		return nil, fmt.Errorf("musicbrainz_api: malformed recording lookup for mbid=%s", mbid)
	}
	if caching.MusicBrainzCacheErrorExists(mbid) {
		log.Debugf("found cached error for %s. Skipping lookup.", mbid)
		d, _, _ := caching.MusicBrainzCacheErrorGet(mbid)
		if d == nil {
			d = fmt.Errorf("cached error is faulty. This most likely is a bug")
		}
		return nil, d
	}
	if caching.MusicBrainzCacheRecordingExists(mbid) {
		log.Infof("Recording with mbid=%s exists in cache", mbid)
		d, _, b := caching.MusicBrainzCacheRecordingGet(mbid)
		dm := d.(musicbrainz_types.Recording)
		dm.MetaInformation = common_types.CreateMetaInformation(true, b)
		return &dm, nil
	}

	log.Infof("Recording with mbid=%s does not exist in cache; looking it up!", mbid)
	apiUrl := fmt.Sprintf("https://musicbrainz.org/ws/2/recording/%s", mbid)
	if includeAliases {
		apiUrl += "?inc=artist-credits+isrcs+releases+works"
	}
	data, err := api_commons.FetchXML[musicbrainz_types.WrappedRecordingLookupResult](apiUrl)
	if err != nil {
		caching.MusicBrainzCacheErrorAdd(mbid, err)
		return nil, err
	}
	data.Recording.MetaInformation = common_types.CreateMetaInformation(true, "")
	caching.MusicBrainzCacheRecordingAdd(mbid, data.Recording)
	return &data.Recording, nil
}
