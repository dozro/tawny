package client

import (
	"github.com/dozro/tawny/pkg/lfm_api"
	"github.com/dozro/tawny/pkg/lfm_types"
)

func GetUserRecentTracks(username string, apikey string, limit int, page int) (*lfm_types.UserGetRecentTracks, error) {
	lt, err := lfm_api.User{}.GetRecentTracks(lfm_api.UserGetArgsWithLimitPage{
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

func GetUserCurrentTrack(username string, apiKey string) (*lfm_types.UserGetRecentTracks, error) {
	return GetUserRecentTracks(username, apiKey, 1, -1)
}
