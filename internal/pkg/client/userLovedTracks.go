package client

import (
	"github.com/dozro/tawny/pkg/lfm_api"
	"github.com/dozro/tawny/pkg/lfm_types"
)

func GetUserLovedTracks(username string, apikey string, limit int, page int) (*lfm_types.UserGetLovedTracks, error) {
	lt, err := lfm_api.User{}.GetLovedTracks(lfm_api.UserGetLovedTracksArgs{
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
