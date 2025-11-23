package lfm_types

type UserGetLovedTracks struct {
	LovedTracks []UserTrack `xml:"track"`
}

type WrappedUserGetLovedTracks struct {
	LovedTracks UserGetLovedTracks `xml:"lovedtracks"`
}
