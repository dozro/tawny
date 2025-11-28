package server_config

import sc "codeberg.org/dozrye/golang_simpleconfig"

type ServerConfig struct {
	ApiPort      int    `json:"api_port"`
	ApiHost      string `json:"api_host"`
	ApiBasePath  string `json:"api_base_path"`
	HmacSecret   string `json:"hmac_secret"`
	LastFMAPIKey string `json:"last_fm_api_key"`
	DebugMode    bool   `json:"debug_mode"`
	ReleaseMode  bool   `json:"release_mode"`
}

func SetupServerConfig() *ServerConfig {
	ch := sc.SimpleConfigHandler{}
	ch.Init("TAWNY", true, true, true, nil)
	apiport := ch.GetIntOption(sc.ConfigEntry{DefaultInt: 8080, Key: "API_PORT", Description: "Port for the TAWNY server"})
	apihost := ch.GetStringOption(sc.ConfigEntry{DefaultString: "localhost", Key: "API_HOST", Description: "Hostname for the TAWNY server"})
	apibasepath := ch.GetStringOption(sc.ConfigEntry{DefaultString: "/api", Key: "API_BASEPATH", Description: "Base path for the TAWNY server"})
	hmacsecret := ch.GetStringOption(sc.ConfigEntry{Key: "HMAC_SECRET", Description: "The HMAC secret for the TAWNY server"})
	lastfmapikey := ch.GetStringOption(sc.ConfigEntry{Key: "LASTFM_API_KEY", Description: "The last fm api key for the TAWNY server"})
	debugmode := ch.GetBooleanOption(sc.ConfigEntry{Key: "DEBUG_MODE", Description: "Debug mode for the TAWNY server", DefaultBool: false})
	releasemode := ch.GetBooleanOption(sc.ConfigEntry{Key: "RELEASE_MODE", Description: "Release mode for the TAWNY server", DefaultBool: true})
	return &ServerConfig{
		ApiPort:      *apiport,
		ApiHost:      *apihost,
		ApiBasePath:  *apibasepath,
		HmacSecret:   *hmacsecret,
		LastFMAPIKey: *lastfmapikey,
		DebugMode:    *debugmode,
		ReleaseMode:  *releasemode,
	}
}
