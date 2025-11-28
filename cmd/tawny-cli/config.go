package main

import (
	"flag"
	"net/url"
)

type tawnyCliConfig struct {
	ApiKey      string
	ApiEndpoint url.URL
	Username    string
	Op          string
}

func Flagread() tawnyCliConfig {
	apiKeyFlag := flag.String("apikey", "", "LastFm API key")
	endpointFlag := flag.String("endpoint", "", "Tawny API endpoint")
	usernameFlag := flag.String("username", "", "Tawny username")
	opFlag := flag.String("op", "user_tracks_current", "Tawny Op")
	flag.Parse()
	return tawnyCliConfig{
		ApiKey: *apiKeyFlag,
		ApiEndpoint: url.URL{
			Scheme: "https",
			Host:   *endpointFlag,
			Path:   "/api/v1",
		},
		Username: *usernameFlag,
		Op:       *opFlag,
	}
}
