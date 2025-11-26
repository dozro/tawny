package lfm_api

import (
	"github.com/dozro/tawny/pkg/api_commons"
	"github.com/dozro/tawny/pkg/lfm_types"
)

func (User) GetTopAlbums(args UserGetArgsWithLimitPage) (*lfm_types.UserGetTopAlbums, error) {
	apiUrl := pageLimitAK(baseUrl, "user.getTopAlbums", args.UserName, args.ApiKey, args.Limit, args.Page)

	data, err := api_commons.FetchXML[lfm_types.WrappedUserGetTopAlbums](apiUrl)

	return &data.UserTopAlbums, err
}
