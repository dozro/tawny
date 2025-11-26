package client

import (
	"github.com/dozro/tawny/pkg/lfm_api"
	"github.com/dozro/tawny/pkg/lfm_types"
	log "github.com/sirupsen/logrus"
)

func GetUserRecentTracks(username string, apikey string, limit int, page int, embedMB bool) (*lfm_types.UserGetRecentTracks, error) {
	log.Debugf("getting recent tracks for %s ...", username)
	lt, err := lfm_api.User{}.GetRecentTracks(lfm_api.UserGetArgsWithLimitPage{
		ApiKey:   apikey,
		UserName: username,
		Limit:    limit,
		Page:     page,
	})
	if embedMB {
		for i, _ := range lt.Track {
			lt.Track[i].EmbedMusicBrainz()
		}
	}
	if err != nil {
		return nil, err
	}
	return lt, nil
}

func GetUserCurrentTrack(username string, apiKey string) (*lfm_types.UserGetRecentTracks, error) {
	return GetUserRecentTracks(username, apiKey, 1, -1, false)
}
