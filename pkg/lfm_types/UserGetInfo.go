package lfm_types

type WrappedUserGetInfo struct {
	User UserGetInfo `xml:"user"`
}

type UserGetInfo struct {
	Id         int64  `xml:"id" json:"id" yaml:"id"`
	Name       string `xml:"name" json:"name" yaml:"name"`
	RealName   string `xml:"realName" json:"real_name" yaml:"real_name"`
	Url        string `xml:"url" json:"url" yaml:"url"`
	ImageUrl   string `xml:"image" json:"image_url" yaml:"image_url"`
	Country    string `xml:"country" json:"country" yaml:"country"`
	Age        int8   `xml:"age" json:"age" yaml:"age"`
	Gender     string `xml:"gender" json:"gender" yaml:"gender"`
	Subscriber int16  `xml:"subscriber" json:"subscriber" yaml:"subscriber"`
	Playcount  int32  `xml:"playcount" json:"playcount" yaml:"playcount"`
	Playlists  int16  `xml:"playlists" json:"playlists" yaml:"playlists"`
	Bootstrap  int16  `xml:"bootstrap" json:"bootstrap" yaml:"bootstrap"`
	Registered string `xml:"registered" json:"registered" yaml:"registered"`
}
