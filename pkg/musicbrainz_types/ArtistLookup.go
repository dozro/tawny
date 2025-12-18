package musicbrainz_types

import (
	"github.com/dozro/tawny/pkg/apiError"
	"github.com/dozro/tawny/pkg/common_types"
)

type WrappedArtistLookupResult struct {
	Artist Artist `xml:"artist" json:"artist,omitempty"`
}

type Artist struct {
	ApiError        apiError.ApiError            `xml:"apiError" json:"api_error,omitempty"`
	Text            string                       `xml:",chardata" json:"text,omitempty"`
	ID              string                       `xml:"id,attr" json:"id,omitempty"`
	Type            string                       `xml:"type,attr" json:"type,omitempty"`
	TypeID          string                       `xml:"type-id,attr" json:"type_id,omitempty"`
	Name            string                       `xml:"name" json:"name,omitempty"`
	SortName        string                       `xml:"sort-name" json:"sort_name"`
	Disambiguation  string                       `xml:"disambiguation"`
	IsniList        IsniList                     `xml:"isni-list" json:"isni_list,omitempty"`
	Country         string                       `xml:"country"`
	Area            Area                         `xml:"area" json:"area,omitempty"`
	BeginArea       Area                         `xml:"begin-area" json:"begin_area,omitempty"`
	LifeSpan        LifeSpan                     `xml:"life-span" json:"life_span,omitempty"`
	AliasList       ArtistAliasList              `xml:"alias-list" json:"alias_list,omitempty"`
	MetaInformation common_types.MetaInformation `xml:"meta-information" json:"meta_information,omitempty"`
}
