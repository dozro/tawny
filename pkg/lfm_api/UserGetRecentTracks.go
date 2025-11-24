package lfm_api

import (
	"github.com/dozro/tawny/pkg/lfm_types"

	log "github.com/sirupsen/logrus"
)

func (User) GetRecentTracks(args UserGetArgsWithLimitPage) (*lfm_types.UserGetRecentTracks, error) {
	apiUrl := pageLimitAK(baseUrl, "user.getRecentTracks", args.UserName, args.ApiKey, args.Limit, args.Page)
	log.Debugf("apiUrl: %s", apiUrl)

	data, err := fetchXML[lfm_types.WrappedUserGetRecentTracks](apiUrl)

	for i := range data.RecentTracks.Track {
		data.RecentTracks.Track[i].Brainz()
	}

	return &data.RecentTracks, err
}
