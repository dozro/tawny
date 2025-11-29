package lfm_types

type WrappedUserGetWeeklyAlbumChart struct {
	WeeklyAlbumChart UserGetWeeklyAlbumChart `xml:"weeklyalbumchart" json:"weekly_album_chart"`
}

type UserGetWeeklyAlbumChart struct {
	User   string     `xml:"user,attr" json:"user"`
	FromTS string     `xml:"from,attr" json:"from_ts"`
	ToTS   string     `xml:"to,attr" json:"to_ts"`
	Album  UserAlbums `xml:"album" json:"album"`
}
