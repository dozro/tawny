package lfm_types

type UserGetTopTracks struct {
	UserAlbums []UserTrack `xml:"track"`
}

type WrappedUserGetTopTracks struct {
	UserTopTracks UserGetTopTracks `xml:"toptracks"`
	Username      string           `xml:"user,attr"`
}
