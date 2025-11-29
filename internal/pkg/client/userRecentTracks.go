package client

import (
	"github.com/dozro/tawny/pkg/lfm_api"
	"github.com/dozro/tawny/pkg/lfm_types"
	log "github.com/sirupsen/logrus"
)

func LfmUserRecentTracks(username, apikey string, limit, page int, embedMB bool) (*lfm_types.UserGetRecentTracks, error) {
	log.Debugf("getting recent tracks for %s ...", username)
	lt, err := lfm_api.User{}.GetRecentTracks(lfm_api.UserGetArgsWithLimitPage{
		ApiKey:   apikey,
		UserName: username,
		Limit:    limit,
		Page:     page,
	})
	if embedMB {
		for i := range lt.Track {
			lt.Track[i].EmbedMusicBrainz()
		}
	}
	if err != nil {
		return nil, err
	}
	return lt, nil
}

func LfmUserCurrentTrack(username, apiKey string, embedMB bool) (*lfm_types.UserGetRecentTracks, error) {
	return LfmUserRecentTracks(username, apiKey, 1, -1, embedMB)
}
