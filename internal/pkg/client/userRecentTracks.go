package client

import (
	"time"

	"github.com/dozro/tawny/pkg/apiError"
	"github.com/dozro/tawny/pkg/lfm_api"
	"github.com/dozro/tawny/pkg/lfm_types"
	"github.com/dozro/tawny/pkg/listenbrainz_api"
	log "github.com/sirupsen/logrus"
)

func LfmUserRecentTracks(username, apikey string, limit, page int, embedMB, embedMBDisabledByServerConfig bool) (*lfm_types.UserGetRecentTracks, error) {
	log.Debugf("getting recent tracks for %s ...", username)
	lt, err := lfm_api.User{}.GetRecentTracks(lfm_api.UserGetArgsWithLimitPage{
		ApiKey:   apikey,
		UserName: username,
		Limit:    limit,
		Page:     page,
	})
	if err != nil {
		return nil, err
	}
	if embedMBDisabledByServerConfig {
		for i := range lt.Track {
			lt.Track[i].SetApiError(apiError.ApiError{
				HttpCode:          503,
				InternalErrorCode: apiError.MusicBrainzLookupDisabledByConfig,
				InternalErrorMsg:  apiError.MusicBrainzLookupDisabledByConfig.String(),
				Message:           "The enrichment of data with MusicBrainz Data is disabled by the server Admin",
				Data:              nil,
				Success:           false,
				Date:              time.Now().String(),
			})
		}
	} else if embedMB {
		for i := range lt.Track {
			lt.Track[i].EmbedMusicBrainz()
		}
	}
	return lt, nil
}

func LbGetCurrentTrack(username string, embedMB, embedMBDisabledByServerConfig bool) (*lfm_types.UserGetRecentTracks, error) {
	log.Debugf("getting recent tracks for %s ...", username)
	lt, err := listenbrainz_api.User{}.GetCurrentTrackLfmCompat(username)
	if err != nil {
		return nil, err
	}
	if embedMBDisabledByServerConfig {
		for i := range lt.Track {
			lt.Track[i].SetApiError(apiError.ApiError{
				HttpCode:          503,
				InternalErrorCode: apiError.MusicBrainzLookupDisabledByConfig,
				InternalErrorMsg:  apiError.MusicBrainzLookupDisabledByConfig.String(),
				Message:           "The enrichment of data with MusicBrainz Data is disabled by the server Admin",
				Data:              nil,
				Success:           false,
				Date:              time.Now().String(),
			})
		}
	} else if embedMB {
		for i := range lt.Track {
			lt.Track[i].EmbedMusicBrainz()
		}
	}
	return lt, nil
}

func LfmUserCurrentTrack(username, apiKey string, embedMB, embedMBDisabledByServerConfig bool) (*lfm_types.UserGetRecentTracks, error) {
	return LfmUserRecentTracks(username, apiKey, 1, -1, embedMB, embedMBDisabledByServerConfig)
}
