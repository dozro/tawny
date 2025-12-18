package musicbrainz_api

import (
	"fmt"

	"github.com/dozro/tawny/internal/pkg/api_commons"
	"github.com/dozro/tawny/pkg/musicbrainz_types"
)

func AreaLookupByMbid(mbid string, includeAliases bool) (*musicbrainz_types.AreaLookupResult, error) {
	if !mbidRegexp.MatchString(mbid) {
		return nil, fmt.Errorf("musicbrainz_api: malformed area lookup")
	}
	apiUrl := fmt.Sprintf("https://musicbrainz.org/ws/2/area/%s", mbid)
	if includeAliases {
		apiUrl += "?inc=aliases"
	}
	data, err := api_commons.FetchXML[musicbrainz_types.AreaLookupResult](apiUrl)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
