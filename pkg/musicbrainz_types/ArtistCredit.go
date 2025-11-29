package musicbrainz_types

type ArtistCredit struct {
	Text       string       `xml:",chardata" json:"text,omitempty"`
	NameCredit []NameCredit `xml:"name-credit" json:"name-credit,omitempty"`
}

type NameCredit struct {
	Text       string           `xml:",chardata" json:"text,omitempty"`
	Joinphrase string           `xml:"joinphrase,attr" json:"joinphrase,omitempty"`
	Artist     NameCreditArtist `xml:"artist" json:"artist,omitempty"`
}

type NameCreditArtist struct {
	Text     string `xml:",chardata" json:"text,omitempty"`
	ID       string `xml:"id,attr" json:"id,omitempty"`
	Type     string `xml:"type,attr" json:"type,omitempty"`
	TypeID   string `xml:"type-id,attr" json:"type-id,omitempty"`
	Name     string `xml:"name"`
	SortName string `xml:"sort-name"`
}
