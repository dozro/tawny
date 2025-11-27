package lfm_types

type UserGetRecentTracks struct {
	Track []LFMTrack `xml:"track"`
}

type WrappedUserGetRecentTracks struct {
	RecentTracks UserGetRecentTracks `xml:"recenttracks"`
}
