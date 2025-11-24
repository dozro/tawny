package lfm_types

import "fmt"

type UserTrack struct {
	Name           string          `xml:"name"`
	Mbid           string          `xml:"mbid"`
	MusicBrainzUrl string          `xml:"musicBrainzUrl"`
	Url            string          `xml:"url"`
	Date           string          `xml:"date"`
	Image          string          `xml:"image"`
	Artist         UserTrackArtist `xml:"artist"`
	Streamable     int8            `xml:"streamable"`
}

type UserTrackArtist struct {
	Name           string `xml:"name"`
	Mbid           string `xml:"mbid"`
	MusicBrainzUrl string `xml:"music_brainz_url"`
	Url            string `xml:"url"`
}

func (ut *UserTrack) Brainz() {
	ut.MusicBrainzUrl = fmt.Sprintf("https://example.org/%s", ut.Mbid)
}
