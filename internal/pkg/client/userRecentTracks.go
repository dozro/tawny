package client

import (
	"lastfm-proxy/pkg/lfm_api"
	"lastfm-proxy/pkg/lfm_types"
)

func GetUserRecentTracks(username string, apikey string, limit int, page int) (*lfm_types.UserGetRecentTracks, error) {
	lt, err := lfm_api.User{}.GetRecentTracks(lfm_api.UserGetLovedTracksArgs{
		ApiKey:   apikey,
		UserName: username,
		Limit:    limit,
		Page:     page,
	})
	if err != nil {
		return nil, err
	}
	return lt, nil
}
