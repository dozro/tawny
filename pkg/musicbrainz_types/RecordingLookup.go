package musicbrainz_types

type WrappedRecordingLookupResult struct {
	Recording Recording `xml:"recording" json:"recording,omitempty"`
}

type Recording struct {
	Text         string               `xml:",chardata" json:"text,omitempty"`
	ID           string               `xml:"id,attr" json:"id,omitempty"`
	Title        string               `xml:"title"`
	Length       int                  `xml:"length"`
	ArtistCredit ArtistCredit         `xml:"artist-credit" json:"artist-credit,omitempty"`
	ReleaseList  RecordingReleaseList `xml:"release-list" json:"release-list,omitempty"`
	IsrcList     IsrcList             `xml:"isrc-list" json:"isrc-list,omitempty"`
}

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

type RecordingReleaseList struct {
	Text    string             `xml:",chardata" json:"text,omitempty"`
	Count   string             `xml:"count,attr" json:"count,omitempty"`
	Release []RecordingRelease `xml:"release" json:"release,omitempty"`
}

type RecordingRelease struct {
	Text  string `xml:",chardata" json:"text,omitempty"`
	ID    string `xml:"id,attr" json:"id,omitempty"`
	Title string `xml:"title"`
}
