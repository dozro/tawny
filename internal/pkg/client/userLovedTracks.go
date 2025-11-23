package client

import (
	"lastfm-proxy/pkg/lfm_api"
	"lastfm-proxy/pkg/lfm_types"
)

func GetUserLovedTracks(username string, apikey string) (*lfm_types.UserGetLovedTracks, error) {
	lt, err := lfm_api.User{}.GetLovedTracks(apikey, username)
	if err != nil {
		return nil, err
	}
	return lt, nil

}
