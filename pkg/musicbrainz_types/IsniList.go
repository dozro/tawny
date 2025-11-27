package musicbrainz_types

type IsniList struct {
	Text string   `xml:",chardata" json:"text,omitempty"`
	Isni []string `xml:"isni"`
}
