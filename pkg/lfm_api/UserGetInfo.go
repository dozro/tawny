package lfm_api

import (
	"fmt"

	"github.com/dozro/tawny/internal/pkg/api_commons"
	"github.com/dozro/tawny/pkg/lfm_types"
)

type UserGetInfoArgs struct {
	ApiKey   string
	UserName string
}

func (User) GetInfo(args UserGetInfoArgs) (*lfm_types.UserGetInfo, error) {
	apiUrl := fmt.Sprintf("%s?method=user.getinfo&user=%s&api_key=%s", baseUrl, args.UserName, args.ApiKey)

	data, err := api_commons.FetchXML[lfm_types.WrappedUserGetInfo](apiUrl)

	return &data.User, err
}
