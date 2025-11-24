package lfm_api

import (
	"github.com/dozro/tawny/pkg/lfm_types"
	log "github.com/sirupsen/logrus"
)

func (User) GetTopAlbums(args UserGetArgsWithLimitPage) (*lfm_types.UserGetTopAlbums, error) {
	apiUrl := pageLimitAK(baseUrl, "user.getTopAlbums", args.UserName, args.ApiKey, args.Limit, args.Page)
	log.Debugf("apiUrl: %s", apiUrl)

	data, err := fetchXML[lfm_types.WrappedUserGetTopAlbums](apiUrl)

	return &data.UserTopAlbums, err
}
