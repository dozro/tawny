package lfm_types

type WrappedUserGetInfo struct {
	User UserGetInfo `xml:"user"`
}

type UserGetInfo struct {
	Id         int64  `xml:"id"`
	Name       string `xml:"name"`
	RealName   string `xml:"realName"`
	Url        string `xml:"url"`
	ImageUrl   string `xml:"image"`
	Country    string `xml:"country"`
	Age        int8   `xml:"age"`
	Gender     string `xml:"gender"`
	Subscriber int16  `xml:"subscriber"`
	Playcount  int32  `xml:"playcount"`
	Playlists  int16  `xml:"playlists"`
	Bootstrap  int16  `xml:"bootstrap"`
	Registered int64  `xml:"registered"`
}
