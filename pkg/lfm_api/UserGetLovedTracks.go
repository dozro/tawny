package lfm_api

import (
	"github.com/dozro/tawny/pkg/api_commons"
	"github.com/dozro/tawny/pkg/lfm_types"
)

type UserGetArgsWithLimitPage struct {
	ApiKey   string
	UserName string
	Limit    int
	Page     int
}

func (User) GetLovedTracks(args UserGetArgsWithLimitPage) (*lfm_types.UserGetLovedTracks, error) {
	apiUrl := pageLimitAK(baseUrl, "user.getLovedTracks", args.UserName, args.ApiKey, args.Limit, args.Page)

	data, err := api_commons.FetchXML[lfm_types.WrappedUserGetLovedTracks](apiUrl)

	return &data.LovedTracks, err
}
