package lfm_types

type UserGetFriends struct {
	Friends UserFriends `xml:"friends"`
}

type UserFriends struct {
	User []UserGetInfo `xml:"user"`
}
