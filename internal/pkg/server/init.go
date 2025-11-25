package server

import "regexp"

var userRegex *regexp.Regexp

func init() {
	userRegex = regexp.MustCompile(`^/user/?$`)
}
