package lfm_types

import "fmt"

// why the fuck ever the xml schema is different for tracks between loved and recent
type UserRTrack struct {
	NowPlaying           bool             `xml:"nowplaying,attr,omitempty"`
	Name                 string           `xml:"name"`
	Album                string           `xml:"album"`
	TrackMbid            string           `xml:"mbid"`
	TrackMusicBrainzUrl  string           `xml:"track_music_brainz_url,attr"`
	Url                  string           `xml:"url"`
	Date                 string           `xml:"date"`
	Timestamp            string           `xml:"uts,attr"`
	Image                string           `xml:"image"`
	Artist               UserRTrackArtist `xml:"artist"`
	ArtistMusicBrainzUrl string           `xml:"artist_music_brainz_url,attr"`
	Streamable           int8             `xml:"streamable"`
}

type UserRTrackArtist struct {
	Mbid string `xml:"mbid,attr"`
	Name string `xml:",chardata"`
}

func (u *UserRTrack) Brainz() {
	u.TrackMusicBrainzUrl = fmt.Sprintf("https://musicbrainz.org/tracks/%s/", u.TrackMbid)
	u.ArtistMusicBrainzUrl = fmt.Sprintf("https://musicbrainz.org/artist/%s/", u.Artist.Mbid)
}

type UserGetRecentTracks struct {
	Track []UserRTrack `xml:"track"`
}

type WrappedUserGetRecentTracks struct {
	RecentTracks UserGetRecentTracks `xml:"recenttracks"`
}
