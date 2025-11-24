package lfm_api

import (
	"fmt"

	"github.com/dozro/tawny/pkg/lfm_types"

	log "github.com/sirupsen/logrus"
)

type UserGetInfoArgs struct {
	ApiKey   string
	UserName string
}

func (User) GetInfo(args UserGetInfoArgs) (*lfm_types.UserGetInfo, error) {
	apiUrl := fmt.Sprintf("%s?method=user.getinfo&user=%s&api_key=%s", baseUrl, args.UserName, args.ApiKey)
	log.Debugf("apiUrl: %s", apiUrl)

	data, err := fetchXML[lfm_types.WrappedUserGetInfo](apiUrl)

	return &data.User, err
}
