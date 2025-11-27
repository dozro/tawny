package musicbrainz_api

import (
	"fmt"

	"github.com/dozro/tawny/internal/pkg/api_commons"
	"github.com/dozro/tawny/pkg/musicbrainz_types"
)

func AreaLookupByMbid(mbid string, includeAliases bool) (*musicbrainz_types.AreaLookupResult, error) {
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

func ArtistLookupByMbid(mbid string, includeAliases bool) (*musicbrainz_types.ArtistLookupResult, error) {
	apiUrl := fmt.Sprintf("https://musicbrainz.org/ws/2/artist/%s", mbid)
	if includeAliases {
		apiUrl += "?inc=aliases"
	}
	data, err := api_commons.FetchXML[musicbrainz_types.ArtistLookupResult](apiUrl)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

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

func RecordingLookupByMbid(mbid string, includeAliases bool) (*musicbrainz_types.RecordingLookupResult, error) {
	apiUrl := fmt.Sprintf("https://musicbrainz.org/ws/2/recording/%s", mbid)
	if includeAliases {
		apiUrl += "?inc=artist-credits+isrcs+releases"
	}
	data, err := api_commons.FetchXML[musicbrainz_types.RecordingLookupResult](apiUrl)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
