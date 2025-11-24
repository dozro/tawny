package lfm_types

import "fmt"

type UserAlbumImage struct {
	Size string `xml:"size,attr"`
	Url  string `xml:",chardata"`
}

type UserAlbums struct {
	Rank           int              `xml:"rank,attr"`
	Name           string           `xml:"name"`
	Playcount      int              `xml:"playcount"`
	Mbid           string           `xml:"mbid"`
	MusicBrainzUrl string           `xml:"music_brainz_url"`
	Url            string           `xml:"url"`
	Artist         UserTrackArtist  `xml:"artist"`
	Image          []UserAlbumImage `xml:"image"`
}

func (ua *UserAlbums) Brainz() {
	ua.MusicBrainzUrl = fmt.Sprintf("https://example.org/%s", ua.Mbid) // To-Do
}

type UserGetTopAlbums struct {
	UserAlbums []UserAlbums `xml:"album"`
}

type WrappedUserGetTopAlbums struct {
	UserTopAlbums UserGetTopAlbums `xml:"topalbums"`
}
