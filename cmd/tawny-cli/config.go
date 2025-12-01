package main

import (
	"flag"
	"net/url"

	sc "codeberg.org/dozrye/golang_simpleconfig"
)

type TawnyCliConfig struct {
	ApiKey      string
	ApiEndpoint url.URL
	Username    string
	HMACSecret  string
	Debug       bool
	Op          string
}

var configHandler sc.SimpleConfigHandler

func Flagread() TawnyCliConfig {
	apiKeyFlag := configHandler.GetStringOption(sc.ConfigEntry{Key: "apikey", Description: "LastFm API Key"})
	endpointFlag := configHandler.GetStringOption(sc.ConfigEntry{Key: "endpoint", DefaultString: "tawny.itsrye.uk", Description: "Tawny API Endpoint"})
	usernameFlag := configHandler.GetStringOption(sc.ConfigEntry{Key: "username", Description: "LastFm Username"})
	hmacSecret := configHandler.GetStringOption(sc.ConfigEntry{Key: "hmacsecret", Description: "Tawny HMAC Secret"})
	debugoutput := configHandler.GetBooleanOption(sc.ConfigEntry{Key: "debug", Description: "Tawny Debug Flag", DefaultBool: false})
	op := configHandler.GetStringOption(sc.ConfigEntry{Key: "op", Description: "Tawny Op"})
	configHandler.ParseFlags()
	flag.Parse()
	return TawnyCliConfig{
		ApiKey: *apiKeyFlag,
		ApiEndpoint: url.URL{
			Scheme: "https",
			Host:   *endpointFlag,
			Path:   "/api/v1",
		},
		Username:   *usernameFlag,
		HMACSecret: *hmacSecret,
		Debug:      *debugoutput,
		Op:         *op,
	}
}
