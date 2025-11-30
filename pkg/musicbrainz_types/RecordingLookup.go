package musicbrainz_types

import "github.com/dozro/tawny/pkg/apiError"

type WrappedRecordingLookupResult struct {
	Recording Recording `xml:"recording" json:"recording,omitempty"`
}

type Recording struct {
	ApiError         apiError.ApiError    `xml:"apiError" json:"api_error,omitempty"`
	Text             string               `xml:",chardata" json:"text,omitempty"`
	ID               string               `xml:"id,attr" json:"id,omitempty"`
	Title            string               `xml:"title"`
	Length           int                  `xml:"length"`
	Disambiguation   string               `xml:"disambiguation" json:"disambiguation,omitempty"`
	FirstReleaseDate string               `xml:"first-release-date" json:"first_release_date,omitempty"`
	ArtistCredit     ArtistCredit         `xml:"artist-credit" json:"artist_credit,omitempty"`
	ReleaseList      RecordingReleaseList `xml:"release-list" json:"release_list,omitempty"`
	IsrcList         IsrcList             `xml:"isrc-list" json:"isrc_list,omitempty"`
}

type RecordingReleaseList struct {
	Text    string    `xml:",chardata" json:"text,omitempty"`
	Count   string    `xml:"count,attr" json:"count,omitempty"`
	Release []Release `xml:"release" json:"release,omitempty"`
}
