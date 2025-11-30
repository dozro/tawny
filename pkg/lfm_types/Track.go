package lfm_types

import (
	"fmt"

	"github.com/dozro/tawny/pkg/apiError"
	"github.com/dozro/tawny/pkg/musicbrainz_api"
	"github.com/dozro/tawny/pkg/musicbrainz_types"
)

type LFMTrack struct {
	Name              string                      `xml:"name" json:"name"`
	Album             string                      `xml:"album" json:"album"`
	Rank              int                         `xml:"rank,attr,omitempty" json:"rank,omitempty"`
	NowPlaying        bool                        `xml:"now_playing,attr,omitempty" json:"now_playing,omitempty"`
	Playcount         int                         `xml:"playcount,omitempty" json:"playcount,omitempty"`
	Mbid              string                      `xml:"mbid" json:"mbid"`
	ArtistMusicBrainz musicbrainz_types.Artist    `xml:"artist_music_brainz,omitempty" json:"artist_music_brainz,omitempty"`
	MusicBrainzUrl    string                      `xml:"musicBrainzUrl,omitempty" json:"musicBrainzUrl,omitempty"`
	TrackMusicBrainz  musicbrainz_types.Recording `xml:"track_music_brainz,omitempty" json:"track_music_brainz,omitempty"`
	Url               string                      `xml:"url" json:"url"`
	Date              string                      `xml:"date,omitempty" json:"date,omitempty"`
	Image             string                      `xml:"image" json:"image"`
	Artist            LFMArtist                   `xml:"artist" json:"artist"`
	Streamable        int8                        `xml:"streamable" json:"streamable"`
}

type LFMArtist struct {
	Name           string `xml:",chardata" json:"name"`
	Mbid           string `xml:"mbid,attr" json:"mbid"`
	MusicBrainzUrl string `xml:"music_brainz_url,omitempty" json:"music_brainz_url,omitempty"`
	Url            string `xml:"url,omitempty" json:"url,omitempty"`
}

func (ut *LFMTrack) String() string {
	return fmt.Sprintf("%s - %s", ut.Name, ut.Artist.Name)
}

func (ut *LFMTrack) Brainz() {
	ut.MusicBrainzUrl = fmt.Sprintf("https://example.org/%s", ut.Mbid)
}

func (t *LFMTrack) SetApiError(apiError apiError.ApiError) {
	t.ArtistMusicBrainz.ApiError = apiError
	t.TrackMusicBrainz.ApiError = apiError
}

func (u *LFMTrack) EmbedMusicBrainz() {
	ma, err := musicbrainz_api.ArtistLookupByMbid(u.Artist.Mbid, false)
	if err == nil {
		u.ArtistMusicBrainz = *ma
	}
	ta, err := musicbrainz_api.RecordingLookupByMbid(u.Mbid, false)
	if err == nil {
		u.TrackMusicBrainz = *ta
	}
}
