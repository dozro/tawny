package musicbrainz_types

type TextRepresentation struct {
	Language string `xml:"language" json:"language,omitempty"`
	Script   string `xml:"script" json:"script,omitempty"`
}
