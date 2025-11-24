package lfm_api

import (
	"fmt"

	"github.com/dozro/tawny/pkg/lfm_types"
	log "github.com/sirupsen/logrus"
)

type UserGetFriendsArgs struct {
	ApiKey       string
	UserName     string
	RecentTracks bool
	Limit        int
	Page         int
}

func (User) GetFriends(args UserGetFriendsArgs) (*lfm_types.UserGetFriends, error) {
	var getRt string = "false"
	if args.RecentTracks {
		getRt = "true"
	}
	apiUrl := fmt.Sprintf("%s&recenttracks=%s", pageLimitAK(baseUrl, "user.getFriends", args.UserName, args.ApiKey, args.Limit, args.Page), getRt)
	log.Debugf("apiUrl: %s", apiUrl)

	data, err := fetchXML[lfm_types.UserGetFriends](apiUrl)

	return &data, err
}
