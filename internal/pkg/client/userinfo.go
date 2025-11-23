package client

import (
	"github.com/dozro/tawny/pkg/lfm_api"
	"github.com/dozro/tawny/pkg/lfm_types"

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
