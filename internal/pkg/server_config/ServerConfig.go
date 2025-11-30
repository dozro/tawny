package server_config

import sc "codeberg.org/dozrye/golang_simpleconfig"

type ServerConfig struct {
	ApiPort                    int                     `json:"api_port"`
	ApiHost                    string                  `json:"api_host"`
	ApiBasePath                string                  `json:"api_base_path"`
	HmacSecret                 string                  `json:"hmac_secret"`
	LastFMAPIKey               string                  `json:"last_fm_api_key"`
	DebugMode                  bool                    `json:"debug_mode"`
	ReleaseMode                bool                    `json:"release_mode"`
	DisabledEndpoints          ServerDisabledEndpoints `json:"disabled_endpoints"`
	DisableEmbeddedMusicBrainz bool                    `json:"disable_embedded_music_brainz"`
}

type ServerDisabledEndpoints struct {
	DisableHMACSigningEndpoint    bool `json:"disable_hmac_signing_endpoint"`
	DisableImageEmbeddedEndpoints bool `json:"disable_image_embedded_endpoint"`
	EnableOnlyHMACEndpoints       bool `json:"enable_only_hmac_endpoints"`
	DisableMusicBrainzEndpoints   bool `json:"disable_music_brainz_endpoints"`
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
	disableHmacSigningEndpoint := ch.GetBooleanOption(sc.ConfigEntry{Key: "DISABLE_HMAC_SIGNING_ENDPOINT", Description: "Disable HMAC Request Signing endpoint", DefaultBool: false})
	disableImageEmbedingEndpoints := ch.GetBooleanOption(sc.ConfigEntry{Key: "DISABLE_IMAGE_EMBEDING_ENDPOINTS", Description: "Disable Image Embeding Endpoints", DefaultBool: false})
	enableOnlyHMACEndpoints := ch.GetBooleanOption(sc.ConfigEntry{Key: "ENABLE_ONLY_HMAC_ENDPOINTS", Description: "Enable only HMAC Request Signing endpoint", DefaultBool: false})
	disableMusicBrainzEndpoints := ch.GetBooleanOption(sc.ConfigEntry{Key: "DISABLE_MUSICBRAINZ_ENDPOINTS", Description: "Disable MusicBrainz Endpoints", DefaultBool: false})
	disableEmbeddedMusicBrainz := ch.GetBooleanOption(sc.ConfigEntry{Key: "DISABLE_MUSICBRAINZ_EMBEDDING", Description: "Disable the enrichment with MusicBrainz data", DefaultBool: false})
	ch.ParseFlags()
	return &ServerConfig{
		ApiPort:                    *apiport,
		ApiHost:                    *apihost,
		ApiBasePath:                *apibasepath,
		HmacSecret:                 *hmacsecret,
		LastFMAPIKey:               *lastfmapikey,
		DebugMode:                  *debugmode,
		ReleaseMode:                *releasemode,
		DisableEmbeddedMusicBrainz: *disableEmbeddedMusicBrainz,
		DisabledEndpoints: ServerDisabledEndpoints{
			DisableHMACSigningEndpoint:    *disableHmacSigningEndpoint,
			DisableMusicBrainzEndpoints:   *disableMusicBrainzEndpoints,
			DisableImageEmbeddedEndpoints: *disableImageEmbedingEndpoints,
			EnableOnlyHMACEndpoints:       *enableOnlyHMACEndpoints,
		},
	}
}
