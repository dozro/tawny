package lfm_api

import (
	"github.com/dozro/tawny/internal/pkg/api_commons"
	"github.com/dozro/tawny/pkg/lfm_types"
	log "github.com/sirupsen/logrus"
)

func (User) GetTopTracks(args UserGetArgsWithLimitPage) (*lfm_types.UserGetTopTracks, error) {
	apiUrl := pageLimitAK(baseUrl, "user.getTopTracks", args.UserName, args.ApiKey, args.Limit, args.Page)

	data, err := api_commons.FetchXML[lfm_types.WrappedUserGetTopTracks](apiUrl)

	if err != nil {
		log.Debugf("Error fetching top tracks: %s", err.Error())
		return nil, err
	}
	return &data.UserTopTracks, err
}
