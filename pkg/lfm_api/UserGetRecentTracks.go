package lfm_api

import (
	"github.com/dozro/tawny/pkg/api_commons"
	"github.com/dozro/tawny/pkg/lfm_types"
)

func (User) GetRecentTracks(args UserGetArgsWithLimitPage) (*lfm_types.UserGetRecentTracks, error) {
	apiUrl := pageLimitAK(baseUrl, "user.getRecentTracks", args.UserName, args.ApiKey, args.Limit, args.Page)

	data, err := api_commons.FetchXML[lfm_types.WrappedUserGetRecentTracks](apiUrl)

	return &data.RecentTracks, err
}
