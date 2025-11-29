package musicbrainz_api

import (
	"fmt"

	"github.com/dozro/tawny/internal/pkg/api_commons"
	"github.com/dozro/tawny/pkg/musicbrainz_types"
)

func RecordingLookupByMbid(mbid string, includeAliases bool) (*musicbrainz_types.Recording, error) {
	if mbid == "" {
		return nil, fmt.Errorf("musicbrainz_api: malformed artist lookup")
	}
	apiUrl := fmt.Sprintf("https://musicbrainz.org/ws/2/recording/%s", mbid)
	if includeAliases {
		apiUrl += "?inc=artist-credits+isrcs+releases"
	}
	data, err := api_commons.FetchXML[musicbrainz_types.WrappedRecordingLookupResult](apiUrl)
	if err != nil {
		return nil, err
	}
	return &data.Recording, nil
}
