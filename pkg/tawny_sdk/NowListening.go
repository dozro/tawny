package tawny_sdk

import (
	"fmt"

	"github.com/dozro/tawny/internal/pkg/api_commons"
	"github.com/dozro/tawny/pkg/lfm_types"
)

func (t Tawny) GetNowListeningFor(username string) (*lfm_types.LFMTrack, error) {
	apiUrl := fmt.Sprintf("%s/user/%s/tracks/current", t.getApiBaseUrlString(), username)
	ct, err := api_commons.FetchJSON[lfm_types.UserGetRecentTracks](apiUrl, t.LastFMApiKey)
	if err != nil {
		return nil, err
	}
	return &ct.Track[0], nil
}
