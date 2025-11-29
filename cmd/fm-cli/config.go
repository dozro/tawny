package main

import (
	"flag"
	"net/url"

	sc "codeberg.org/dozrye/golang_simpleconfig"
)

type tawnyCliConfig struct {
	ApiKey      string
	ApiEndpoint url.URL
	Username    string
	Op          string
}

var configHandler sc.SimpleConfigHandler

func Flagread() tawnyCliConfig {
	apiKeyFlag := configHandler.GetStringOption(sc.ConfigEntry{Key: "apikey", Description: "LastFm API Key"})
	endpointFlag := configHandler.GetStringOption(sc.ConfigEntry{Key: "endpoint", DefaultString: "tawny.itsrye.uk", Description: "Tawny API Endpoint"})
	usernameFlag := configHandler.GetStringOption(sc.ConfigEntry{Key: "username", Description: "LastFm Username"})
	opFlag := configHandler.GetStringOption(sc.ConfigEntry{Key: "op", DefaultString: "user_tracks_current", Description: "Tawny Operation"})
	configHandler.ParseFlags()
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
