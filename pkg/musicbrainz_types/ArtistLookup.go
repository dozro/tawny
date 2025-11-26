package musicbrainz_types

type ArtistLookupResult struct {
	Artist struct {
		Text           string `xml:",chardata" json:"text,omitempty"`
		ID             string `xml:"id,attr" json:"id,omitempty"`
		Type           string `xml:"type,attr" json:"type,omitempty"`
		TypeID         string `xml:"type-id,attr" json:"type-id,omitempty"`
		Name           string `xml:"name"`
		SortName       string `xml:"sort-name"`
		Disambiguation string `xml:"disambiguation"`
		IsniList       struct {
			Text string   `xml:",chardata" json:"text,omitempty"`
			Isni []string `xml:"isni"`
		} `xml:"isni-list" json:"isni-list,omitempty"`
		Country string `xml:"country"`
		Area    struct {
			Text             string `xml:",chardata" json:"text,omitempty"`
			ID               string `xml:"id,attr" json:"id,omitempty"`
			Name             string `xml:"name"`
			SortName         string `xml:"sort-name"`
			Iso31661CodeList struct {
				Text         string `xml:",chardata" json:"text,omitempty"`
				Iso31661Code string `xml:"iso-3166-1-code"`
			} `xml:"iso-3166-1-code-list" json:"iso-3166-1-code-list,omitempty"`
		} `xml:"area" json:"area,omitempty"`
		BeginArea struct {
			Text     string `xml:",chardata" json:"text,omitempty"`
			ID       string `xml:"id,attr" json:"id,omitempty"`
			Name     string `xml:"name"`
			SortName string `xml:"sort-name"`
		} `xml:"begin-area" json:"begin-area,omitempty"`
		LifeSpan struct {
			Text  string `xml:",chardata" json:"text,omitempty"`
			Begin string `xml:"begin"`
			End   string `xml:"end"`
			Ended string `xml:"ended"`
		} `xml:"life-span" json:"life-span,omitempty"`
		AliasList struct {
			Text  string `xml:",chardata" json:"text,omitempty"`
			Count string `xml:"count,attr" json:"count,omitempty"`
			Alias []struct {
				Text     string `xml:",chardata" json:"text,omitempty"`
				SortName string `xml:"sort-name,attr" json:"sort-name,omitempty"`
				Locale   string `xml:"locale,attr" json:"locale,omitempty"`
				Type     string `xml:"type,attr" json:"type,omitempty"`
				TypeID   string `xml:"type-id,attr" json:"type-id,omitempty"`
				Primary  string `xml:"primary,attr" json:"primary,omitempty"`
			} `xml:"alias" json:"alias,omitempty"`
		} `xml:"alias-list" json:"alias-list,omitempty"`
	} `xml:"artist" json:"artist,omitempty"`
}
