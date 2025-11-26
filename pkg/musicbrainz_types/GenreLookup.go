package musicbrainz_types

type GenreLookupResult struct {
	Genre struct {
		Text string `xml:",chardata" json:"text,omitempty"`
		ID   string `xml:"id,attr" json:"id,omitempty"`
		Name string `xml:"name"`
	} `xml:"genre" json:"genre,omitempty"`
}
