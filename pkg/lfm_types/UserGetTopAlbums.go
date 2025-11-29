package lfm_types

import "fmt"

type UserAlbumImage struct {
	Size string `xml:"size,attr"`
	Url  string `xml:",chardata"`
}

type UserAlbums struct {
	Rank           int              `xml:"rank,attr" json:"rank,omitempty"`
	Name           string           `xml:"name"`
	Playcount      int              `xml:"playcount" json:"playcount,omitempty"`
	Mbid           string           `xml:"mbid" json:"mbid,omitempty"`
	MusicBrainzUrl string           `xml:"music_brainz_url" json:"music_brainz_url,omitempty"`
	Url            string           `xml:"url"`
	Artist         LFMArtist        `xml:"artist"`
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
