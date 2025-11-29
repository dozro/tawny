package musicbrainz_api

import (
	"fmt"

	"github.com/dozro/tawny/internal/pkg/api_commons"
	"github.com/dozro/tawny/pkg/musicbrainz_types"
)

func GenreLookupByMbid(mbid string, includeAliases bool) (*musicbrainz_types.GenreLookupResult, error) {
	apiUrl := fmt.Sprintf("https://musicbrainz.org/ws/2/genre/%s", mbid)
	if includeAliases {
		apiUrl += "?inc=aliases"
	}
	data, err := api_commons.FetchXML[musicbrainz_types.GenreLookupResult](apiUrl)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
