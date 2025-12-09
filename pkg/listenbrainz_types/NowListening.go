package listenbrainz_types

import (
	"fmt"

	"github.com/dozro/tawny/pkg/lfm_types"
	"github.com/dozro/tawny/pkg/musicbrainz_types"
	log "github.com/sirupsen/logrus"
)

type NowListeningWrapperOuter struct {
	Payload NowListeningWrapperInner `json:"payload"`
}

type NowListeningWrapperInner struct {
	Count      int      `json:"count"`
	Listens    []Listen `json:"listens"`
	PlayingNow bool     `json:"playing_now"`
	UserId     string   `json:"user_id"`
}

type Listen struct {
	PlayingNow    bool          `json:"playing_now"`
	TrackMetadata TrackMetaData `json:"track_metadata"`
}

func (n NowListeningWrapperInner) ConvertToLfm() (*lfm_types.UserGetRecentTracks, error) {
	if len(n.Listens) == 0 {
		return nil, fmt.Errorf("no listens fetched")
	}
	log.Debug(n)
	lfmTrack := lfm_types.LFMTrack{
		Name:              n.Listens[0].TrackMetadata.TrackName,
		Album:             n.Listens[0].TrackMetadata.ReleaseName,
		Rank:              0,
		NowPlaying:        n.Listens[0].PlayingNow,
		Playcount:         0,
		Mbid:              n.Listens[0].TrackMetadata.TrackMbid,
		ArtistMusicBrainz: musicbrainz_types.Artist{},
		TrackMusicBrainz:  musicbrainz_types.Recording{},
		Url:               "",
		Date:              "",
		Image:             "",
		Artist: lfm_types.LFMArtist{
			Name:           n.Listens[0].TrackMetadata.ArtistName,
			Mbid:           "",
			MusicBrainzUrl: "",
			Url:            "",
		},
		Streamable: 0,
		DataSource: "listenbrainz.org",
		Compat: lfm_types.LFMTrackCompat{
			CompatDate:         "2025-12-09",
			CompatMode:         true,
			CompatRating:       "partial",
			CompatSourceFormat: "ListenBrainz",
		},
	}
	var tracks []lfm_types.LFMTrack
	tracks = append(tracks, lfmTrack)
	return &lfm_types.UserGetRecentTracks{
		Track: tracks,
	}, nil
}
