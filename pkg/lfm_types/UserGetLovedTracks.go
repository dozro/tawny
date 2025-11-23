package lfm_types

type LovedTrackArtist struct {
	Name string `xml:"name,attr"`
	Mbid string `xml:"mbid,attr"`
	Url  string `xml:"url,attr"`
}

type LovedTrack struct {
	Name  string `xml:"name"`
	Mbid  string `xml:"mbid"`
	Url   string `xml:"url"`
	Date  string `xml:"date"`
	Image string `xml:"image"`
}
type UserGetLovedTracks struct {
	LovedTracks []LovedTrack `xml:"track"`
}

type WrappedUserGetLovedTracks struct {
	LovedTracks UserGetLovedTracks `xml:"lovedtracks"`
}
