package lfm_api

import (
	"github.com/dozro/tawny/internal/pkg/api_commons"
	"github.com/dozro/tawny/pkg/lfm_types"
	log "github.com/sirupsen/logrus"
)

func (User) GetWeeklyChart(args FromToAKArgs) (*lfm_types.UserGetWeeklyAlbumChart, error) {
	apiUrl := fromToAK(baseUrl, "user.getWeeklyChart", args)
	data, err := api_commons.FetchXML[lfm_types.WrappedUserGetWeeklyAlbumChart](apiUrl)

	if err != nil {
		log.Debugf("Error fetching weekly chart: %s", err.Error())
	}

	return &data.WeeklyAlbumChart, err
}
