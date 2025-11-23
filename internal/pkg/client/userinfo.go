package client

import (
	"lastfm-proxy/pkg/lfm_api"
	"lastfm-proxy/pkg/lfm_types"

	log "github.com/sirupsen/logrus"
)

func GetUserInfo(username string, apikey string) (*lfm_types.UserGetInfo, error) {
	log.Info("getUserInfo: ", username, apikey)
	userinfo, err := lfm_api.User{}.GetInfo(lfm_api.UserGetInfoArgs{
		UserName: username,
		ApiKey:   apikey,
	})
	if err != nil {
		return nil, err
	}
	return userinfo, nil

}
