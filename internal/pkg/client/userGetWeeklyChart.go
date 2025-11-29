package client

import (
	"github.com/dozro/tawny/pkg/lfm_api"
	"github.com/dozro/tawny/pkg/lfm_types"
	log "github.com/sirupsen/logrus"
)

func GetUserWeeklyChart(username string, apikey string, from int, to int) (*lfm_types.UserGetWeeklyAlbumChart, error) {
	log.Debugf("getting top tracks for %s ...", username)
	wac, err := lfm_api.User{}.GetWeeklyChart(lfm_api.FromToAKArgs{
		ApiKey:   apikey,
		UserName: username,
		From:     from,
		To:       to,
	})
	if err != nil {
		return nil, err
	}
	return wac, nil
}
