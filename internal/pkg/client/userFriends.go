package client

import (
	"github.com/dozro/tawny/pkg/lfm_api"
	"github.com/dozro/tawny/pkg/lfm_types"
)

func LfmUserFriends(username, apikey string, limit, page int, recent bool) (*lfm_types.UserGetFriends, error) {
	uf, err := lfm_api.User{}.GetFriends(lfm_api.UserGetFriendsArgs{
		ApiKey:       apikey,
		UserName:     username,
		RecentTracks: recent,
		Limit:        limit,
		Page:         page,
	})
	if err != nil {
		return nil, err
	}
	return uf, nil
}
