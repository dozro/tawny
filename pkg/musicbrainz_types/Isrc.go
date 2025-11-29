package musicbrainz_types

type IsrcList struct {
	Count string `xml:"count,attr" json:"count,omitempty"`
	Isrc  Isrc   `xml:"isrc" json:"isrc,omitempty"`
}

type Isrc struct {
	Text string `xml:",chardata" json:"text,omitempty"`
	ID   string `xml:"id,attr" json:"id,omitempty"`
}
