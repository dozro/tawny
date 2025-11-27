package lfm_types

type UserGetLovedTracks struct {
	LovedTracks []LFMTrack `xml:"track"`
}

type WrappedUserGetLovedTracks struct {
	LovedTracks UserGetLovedTracks `xml:"lovedtracks"`
}
