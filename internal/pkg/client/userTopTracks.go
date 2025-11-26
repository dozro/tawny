package client

import (
	"github.com/dozro/tawny/pkg/lfm_api"
	"github.com/dozro/tawny/pkg/lfm_types"
	log "github.com/sirupsen/logrus"
)

func GetUserTopTracks(username string, apikey string, limit int, page int) (*lfm_types.UserGetTopTracks, error) {
	log.Debugf("getting top tracls for %s ...", username)
	lt, err := lfm_api.User{}.GetTopTracks(lfm_api.UserGetArgsWithLimitPage{
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
