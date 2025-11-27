package server

import "regexp"

var hmacProxyUserInfoRegex *regexp.Regexp
var hmacProxyUserNowPlayingRegex *regexp.Regexp
var hmacProxyUserNowPlayingEmbed *regexp.Regexp
var hmacProxyUserRecentlyPlayedRegex *regexp.Regexp

func init() {
	hmacProxyUserInfoRegex = regexp.MustCompile(`^[/_]?user/?$`)
	hmacProxyUserNowPlayingRegex = regexp.MustCompile(`^[/_]?user[/_]tracks[/_]current/?$`)
	hmacProxyUserNowPlayingEmbed = regexp.MustCompile(`^^[/_]?user[/_]tracks[/_]current[/_]embed/?$`)
	hmacProxyUserRecentlyPlayedRegex = regexp.MustCompile(`^[/_]?user[/_]tracks[/_]recent/?$`)
}
