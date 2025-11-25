package server

import "regexp"

var userInfoRegex *regexp.Regexp
var userNowPlayingRegex *regexp.Regexp
var userNowPlayingEmbed *regexp.Regexp

func init() {
	userInfoRegex = regexp.MustCompile(`^/user/?$`)
	userNowPlayingRegex = regexp.MustCompile(`^/user/tracks/current/?$`)
	userNowPlayingEmbed = regexp.MustCompile(`^/user/tracks/current/embed?$`)
}
