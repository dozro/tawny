package lfm_types

type UserTrack struct {
	Name       string          `xml:"name"`
	Mbid       string          `xml:"mbid"`
	Url        string          `xml:"url"`
	Date       string          `xml:"date"`
	Image      string          `xml:"image"`
	Artist     UserTrackArtist `xml:"artist"`
	Streamable int8            `xml:"streamable"`
}

type UserTrackArtist struct {
	Name string `xml:"name"`
	Mbid string `xml:"mbid"`
	Url  string `xml:"url"`
}
