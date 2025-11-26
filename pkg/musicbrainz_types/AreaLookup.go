package musicbrainz_types

type AreaLookupResult struct {
	Area struct {
		Text      string `xml:",chardata" json:"text,omitempty"`
		ID        string `xml:"id,attr" json:"id,omitempty"`
		Type      string `xml:"type,attr" json:"type,omitempty"`
		TypeID    string `xml:"type-id,attr" json:"type-id,omitempty"`
		Name      string `xml:"name"`
		SortName  string `xml:"sort-name"`
		AliasList struct {
			Text  string `xml:",chardata" json:"text,omitempty"`
			Count string `xml:"count,attr" json:"count,omitempty"`
			Alias []struct {
				Text     string `xml:",chardata" json:"text,omitempty"`
				Locale   string `xml:"locale,attr" json:"locale,omitempty"`
				SortName string `xml:"sort-name,attr" json:"sort-name,omitempty"`
				Type     string `xml:"type,attr" json:"type,omitempty"`
				TypeID   string `xml:"type-id,attr" json:"type-id,omitempty"`
				Primary  string `xml:"primary,attr" json:"primary,omitempty"`
			} `xml:"alias" json:"alias,omitempty"`
		} `xml:"alias-list" json:"alias-list,omitempty"`
	} `xml:"area" json:"area,omitempty"`
}
