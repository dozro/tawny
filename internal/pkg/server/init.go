package server

import "regexp"

var hmacProxyUserInfoRegex *regexp.Regexp
var hmacProxyUserNowPlayingRegex *regexp.Regexp
var hmacProxyUserNowPlayingEmbed *regexp.Regexp
var hmacProxyUserRecentlyPlayedRegex *regexp.Regexp

var middlewareHMACEndpointRegex *regexp.Regexp
var middlewareEmbedEndpointRegex *regexp.Regexp
var middlewareHMACSignEndpointRegex *regexp.Regexp
var middlewareMusicBrainzEndpointRegex *regexp.Regexp
var supportedImageTypes *regexp.Regexp

func init() {
	hmacProxyUserInfoRegex = regexp.MustCompile(`^[/_]?user/?$`)
	hmacProxyUserNowPlayingRegex = regexp.MustCompile(`^[/_]?user[/_]tracks[/_]current/?$`)
	hmacProxyUserNowPlayingEmbed = regexp.MustCompile(`^^[/_]?user[/_]tracks[/_]current[/_]embed/?$`)
	hmacProxyUserRecentlyPlayedRegex = regexp.MustCompile(`^[/_]?user[/_]tracks[/_]recent/?$`)
	middlewareHMACEndpointRegex = regexp.MustCompile(`^/?hmac/`)
	middlewareEmbedEndpointRegex = regexp.MustCompile(`/embed$`)
	middlewareHMACSignEndpointRegex = regexp.MustCompile(`^/?hmac/sign`)
	middlewareMusicBrainzEndpointRegex = regexp.MustCompile(`^/?musicbrainz/`)
	supportedImageTypes = regexp.MustCompile(`image/(png|tiff|jpeg)`)
}
