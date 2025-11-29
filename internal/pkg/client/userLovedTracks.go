package client

import (
	"github.com/dozro/tawny/pkg/lfm_api"
	"github.com/dozro/tawny/pkg/lfm_types"
	log "github.com/sirupsen/logrus"
)

func UserLovedTracks(username, apikey string, limit, page int) (*lfm_types.UserGetLovedTracks, error) {
	log.Debugf("getting loved tracks for %s ...", username)
	lt, err := lfm_api.User{}.GetLovedTracks(lfm_api.UserGetArgsWithLimitPage{
		ApiKey:   apikey,
		UserName: username,
		Limit:    limit,
		Page:     page,
	})
	if err != nil {
		return nil, err
	}
	return lt, nil
}
