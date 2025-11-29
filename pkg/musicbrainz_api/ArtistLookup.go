package musicbrainz_api

import (
	"fmt"

	"github.com/dozro/tawny/internal/pkg/api_commons"
	"github.com/dozro/tawny/pkg/musicbrainz_types"
)

func ArtistLookupByMbid(mbid string, includeAliases bool) (*musicbrainz_types.Artist, error) {
	if mbid == "" {
		return nil, fmt.Errorf("musicbrainz_api: malformed artist lookup")
	}
	apiUrl := fmt.Sprintf("https://musicbrainz.org/ws/2/artist/%s", mbid)
	if includeAliases {
		apiUrl += "?inc=aliases"
	}
	data, err := api_commons.FetchXML[musicbrainz_types.WrappedArtistLookupResult](apiUrl)
	if err != nil {
		return nil, err
	}
	return &data.Artist, nil
}
