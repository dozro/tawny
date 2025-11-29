package musicbrainz_types

type ReleaseEvent struct {
	Date string `xml:"date" json:"date"`
	Area Area   `xml:"area" json:"area"`
}
