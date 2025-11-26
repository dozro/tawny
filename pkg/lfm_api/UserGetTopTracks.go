package lfm_api

import "github.com/dozro/tawny/pkg/lfm_types"

func (User) GetTopTracks(args UserGetArgsWithLimitPage) (*lfm_types.UserGetTopTracks, error) {
	apiUrl := pageLimitAK(baseUrl, "user.getTopTracks", args.UserName, args.ApiKey, args.Limit, args.Page)

	data, err := fetchXML[lfm_types.WrappedUserGetTopTracks](apiUrl)

	return &data.UserTopTracks, err
}
