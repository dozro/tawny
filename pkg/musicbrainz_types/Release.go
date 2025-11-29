package musicbrainz_types

type Release struct {
	ID                 string             `xml:"id,attr" json:"id"`
	Title              string             `xml:"title" json:"title"`
	Status             ReleaseStatus      `xml:"status" json:"status"`
	Quality            string             `xml:"quality" json:"quality"`
	TextRepresentation TextRepresentation `xml:"text-representation" json:"text_representation"`
	ArtistCredit       ArtistCredit       `xml:"artist-credit" json:"artist_credit"`
	Date               string             `xml:"date" json:"date"`
	Country            string             `xml:"country" json:"country"`
}

type ReleaseStatus struct {
	ID     string `xml:"id,attr" json:"id,omitempty"`
	Status string `xml:",chardata" json:"status,omitempty"`
}
