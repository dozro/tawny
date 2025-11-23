package lfm_types

// why the fuck ever the xml schema is different for tracks between loved and recent
type UserRTrack struct {
	NowPlaying bool   `xml:"nowplaying,attr,omitempty"`
	Name       string `xml:"name"`
	TrackMbid  string `xml:"mbid"`
	Url        string `xml:"url"`
	Date       string `xml:"date"`
	Timestamp  string `xml:"uts,attr"`
	Image      string `xml:"image"`
	Artist     string `xml:"artist"`
	ArtistMbid string `xml:"mbid,attr"`
	Streamable int8   `xml:"streamable"`
}
type UserGetRecentTracks struct {
	Track []UserRTrack `xml:"track"`
}

type WrappedUserGetRecentTracks struct {
	RecentTracks UserGetRecentTracks `xml:"recenttracks"`
}
