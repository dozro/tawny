package musicbrainz_types

type Area struct {
	Text      string        `xml:",chardata" json:"text,omitempty"`
	ID        string        `xml:"id,attr" json:"id,omitempty"`
	Type      string        `xml:"type,attr" json:"type,omitempty"`
	TypeID    string        `xml:"type-id,attr" json:"type_id,omitempty"`
	Name      string        `xml:"name" json:"name"`
	SortName  string        `xml:"sort-name" json:"sort_name"`
	AliasList AreaAliasList `xml:"alias-list" json:"alias_list,omitempty"`
}

type AreaLookupResult struct {
	Area Area `xml:"area" json:"area,omitempty"`
}
