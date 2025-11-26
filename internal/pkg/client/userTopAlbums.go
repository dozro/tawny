package client

import (
	"github.com/dozro/tawny/pkg/lfm_api"
	"github.com/dozro/tawny/pkg/lfm_types"
	log "github.com/sirupsen/logrus"
)

func GetUserTopAlbum(username string, apikey string, limit int, page int) (*lfm_types.UserGetTopAlbums, error) {
	log.Debugf("getting top albums for %s ...", username)
	lt, err := lfm_api.User{}.GetTopAlbums(lfm_api.UserGetArgsWithLimitPage{
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
