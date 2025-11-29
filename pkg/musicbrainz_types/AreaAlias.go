package musicbrainz_types

type AreaAliasList struct {
	Count string      `xml:"count,attr" json:"count,omitempty"`
	Alias []AreaAlias `xml:"alias" json:"alias,omitempty"`
}

type AreaAlias struct {
	Locale   string `xml:"locale,attr" json:"locale,omitempty"`
	SortName string `xml:"sort-name,attr" json:"sort_name,omitempty"`
	Type     string `xml:"type,attr" json:"type,omitempty"`
	TypeID   string `xml:"type-id,attr" json:"type_id,omitempty"`
	Primary  string `xml:"primary,attr" json:"primary,omitempty"`
}
