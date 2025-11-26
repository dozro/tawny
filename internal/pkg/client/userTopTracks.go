package client

import (
	"github.com/dozro/tawny/pkg/lfm_api"
	"github.com/dozro/tawny/pkg/lfm_types"
)

func GetUserTopTracks(username string, apikey string, limit int, page int) (*lfm_types.UserGetTopTracks, error) {
	lt, err := lfm_api.User{}.GetTopTracks(lfm_api.UserGetArgsWithLimitPage{
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
