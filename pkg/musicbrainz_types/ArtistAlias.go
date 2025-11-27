package musicbrainz_types

type ArtistAliasList struct {
	Text  string        `xml:",chardata" json:"text,omitempty"`
	Count string        `xml:"count,attr" json:"count,omitempty"`
	Alias []ArtistAlias `xml:"alias" json:"alias,omitempty"`
}

type ArtistAlias struct {
	Text     string `xml:",chardata" json:"text,omitempty"`
	SortName string `xml:"sort-name,attr" json:"sort-name,omitempty"`
	Locale   string `xml:"locale,attr" json:"locale,omitempty"`
	Type     string `xml:"type,attr" json:"type,omitempty"`
	TypeID   string `xml:"type-id,attr" json:"type-id,omitempty"`
	Primary  string `xml:"primary,attr" json:"primary,omitempty"`
}
