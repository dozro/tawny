package lfm_types

import (
	"github.com/dozro/tawny/pkg/musicbrainz_api"
	"github.com/dozro/tawny/pkg/musicbrainz_types"
)

// why the fuck ever the xml schema is different for tracks between loved and recent
type UserRTrack struct {
	NowPlaying        bool                                    `xml:"nowplaying,attr,omitempty"`
	Name              string                                  `xml:"name"`
	Album             string                                  `xml:"album"`
	TrackMbid         string                                  `xml:"mbid"`
	TrackMusicBrainz  musicbrainz_types.RecordingLookupResult `xml:"track_music_brainz,omitempty"`
	Url               string                                  `xml:"url"`
	Date              string                                  `xml:"date"`
	Timestamp         string                                  `xml:"uts,attr"`
	Image             string                                  `xml:"image"`
	Artist            UserRTrackArtist                        `xml:"artist"`
	ArtistMusicBrainz musicbrainz_types.ArtistLookupResult    `xml:"artist_music_brainz,omitempty" json:"artist_music_brainz,omitempty"`
	Streamable        int8                                    `xml:"streamable"`
}

type UserRTrackArtist struct {
	Mbid string `xml:"mbid,attr"`
	Name string `xml:",chardata"`
}

func (u *UserRTrack) EmbedMusicBrainz() {
	ma, err := musicbrainz_api.ArtistLookupByMbid(u.Artist.Mbid, false)
	if err == nil {
		u.ArtistMusicBrainz = *ma
	}
	ta, err := musicbrainz_api.RecordingLookupByMbid(u.TrackMbid, false)
	if err == nil {
		u.TrackMusicBrainz = *ta
	}
}

type UserGetRecentTracks struct {
	Track []UserRTrack `xml:"track"`
}

type WrappedUserGetRecentTracks struct {
	RecentTracks UserGetRecentTracks `xml:"recenttracks"`
}
