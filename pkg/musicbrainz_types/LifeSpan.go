package musicbrainz_types

type LifeSpan struct {
	Text  string `xml:",chardata" json:"text,omitempty"`
	Begin string `xml:"begin"`
	End   string `xml:"end"`
	Ended string `xml:"ended"`
}
