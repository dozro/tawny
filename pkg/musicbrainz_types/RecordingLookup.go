package musicbrainz_types

type RecordingLookupResult struct {
	Recording struct {
		Text         string `xml:",chardata" json:"text,omitempty"`
		ID           string `xml:"id,attr" json:"id,omitempty"`
		Title        string `xml:"title"`
		Length       int    `xml:"length"`
		ArtistCredit struct {
			Text       string `xml:",chardata" json:"text,omitempty"`
			NameCredit []struct {
				Text       string `xml:",chardata" json:"text,omitempty"`
				Joinphrase string `xml:"joinphrase,attr" json:"joinphrase,omitempty"`
				Artist     struct {
					Text     string `xml:",chardata" json:"text,omitempty"`
					ID       string `xml:"id,attr" json:"id,omitempty"`
					Type     string `xml:"type,attr" json:"type,omitempty"`
					TypeID   string `xml:"type-id,attr" json:"type-id,omitempty"`
					Name     string `xml:"name"`
					SortName string `xml:"sort-name"`
				} `xml:"artist" json:"artist,omitempty"`
			} `xml:"name-credit" json:"name-credit,omitempty"`
		} `xml:"artist-credit" json:"artist-credit,omitempty"`
		ReleaseList struct {
			Text    string `xml:",chardata" json:"text,omitempty"`
			Count   string `xml:"count,attr" json:"count,omitempty"`
			Release []struct {
				Text  string `xml:",chardata" json:"text,omitempty"`
				ID    string `xml:"id,attr" json:"id,omitempty"`
				Title string `xml:"title"`
			} `xml:"release" json:"release,omitempty"`
		} `xml:"release-list" json:"release-list,omitempty"`
		IsrcList struct {
			Text  string `xml:",chardata" json:"text,omitempty"`
			Count string `xml:"count,attr" json:"count,omitempty"`
			Isrc  struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				ID   string `xml:"id,attr" json:"id,omitempty"`
			} `xml:"isrc" json:"isrc,omitempty"`
		} `xml:"isrc-list" json:"isrc-list,omitempty"`
	} `xml:"recording" json:"recording,omitempty"`
}
