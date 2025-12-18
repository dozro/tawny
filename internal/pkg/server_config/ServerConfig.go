package server_config

import (
	"os"

	sc "codeberg.org/dozrye/golang_simpleconfig"
)

type ServerConfig struct {
	ApiPort              int                     `json:"api_port"`
	ApiHost              string                  `json:"api_host"`
	ApiBasePath          string                  `json:"api_base_path"`
	HmacSecret           string                  `json:"hmac_secret"`
	LastFMAPIKey         string                  `json:"last_fm_api_key"`
	DebugMode            bool                    `json:"debug_mode"`
	ReleaseMode          bool                    `json:"release_mode"`
	DevelopMode          bool                    `json:"develop_mode"`
	CachingConfig        CachingConfig           `json:"redis_config"`
	DisabledEndpoints    ServerDisabledEndpoints `json:"disabled_endpoints"`
	ExtendedServerConfig ExtendedServerConfig    `json:"extended_server_config"`
	ServerOperator       ServerOperatorInfo      `json:"server_operator"`
}

type CachingConfig struct {
	EnableRedis    bool   `json:"enable_redis"`
	DisableCaching bool   `json:"disable_caching"`
	LocalCachePath string `json:"local_cache_path"`
	SaveCacheToFS  bool   `json:"save_cache_to_fs"`
}

type ServerDisabledEndpoints struct {
	DisableHMACSigningEndpoint    bool `json:"disable_hmac_signing_endpoint"`
	DisableImageEmbeddedEndpoints bool `json:"disable_image_embedded_endpoint"`
	EnableOnlyHMACEndpoints       bool `json:"enable_only_hmac_endpoints"`
	DisableMusicBrainzEndpoints   bool `json:"disable_music_brainz_endpoints"`
	DisableSwaggerUI              bool `json:"disable_swagger_ui"`
}

type ServerOperatorInfo struct {
	OperatorName     string `json:"operator_name"`
	OperatorContact  string `json:"operator_contact"`
	PrivacyPolicyURL string `json:"privacy_policy_url"`
	ImprintURL       string `json:"imprint_url"`
}

type ExtendedServerConfig struct {
	RunningInDocker            bool   `json:"running_in_docker"`
	LogOutputFormat            string `json:"log_output_format"`
	DisableEmbeddedMusicBrainz bool   `json:"disable_embedded_music_brainz"`
	DisableGinVersionPublished bool   `json:"disable_gin_version_published"`
	HideSensitiveInformation   bool   `json:"hide_sensitive_information"`
	TawnyVersion               string `json:"tawny_version"`
	TawnyRevision              string `json:"tawny_revision"`
	SourceCodeURL              string `json:"source_code_url"`
	SourceAbuseContact         string `json:"source_abuse_contact"`
}

func SetupServerConfig() *ServerConfig {
	ch := sc.SimpleConfigHandler{}
	ch.Init("TAWNY", true, true, true, nil)
	apiport := ch.GetIntOption(sc.ConfigEntry{DefaultInt: 8080, Key: "API_PORT", Description: "Port for the TAWNY server"})
	hostname, _ := os.Hostname()
	apihost := ch.GetStringOption(sc.ConfigEntry{DefaultString: hostname, Key: "API_HOST", Description: "Hostname for the TAWNY server"})
	apibasepath := ch.GetStringOption(sc.ConfigEntry{DefaultString: "/api", Key: "API_BASEPATH", Description: "Base path for the TAWNY server"})
	hmacsecret := ch.GetStringOption(sc.ConfigEntry{Key: "HMAC_SECRET", Description: "The HMAC secret for the TAWNY server"})
	lastfmapikey := ch.GetStringOption(sc.ConfigEntry{Key: "LASTFM_API_KEY", Description: "The last fm api key for the TAWNY server"})
	debugmode := ch.GetBooleanOption(sc.ConfigEntry{Key: "DEBUG_MODE", Description: "Debug mode for the TAWNY server", DefaultBool: false})
	releasemode := ch.GetBooleanOption(sc.ConfigEntry{Key: "RELEASE_MODE", Description: "Release mode for the TAWNY server", DefaultBool: true})
	developmode := ch.GetBooleanOption(sc.ConfigEntry{Key: "DEVELOP_MODE", Description: "Develop mode for the TAWNY server, THIS CHANGES SERVER BEHAVIOUR", DefaultBool: false})
	runningInDocker := ch.GetBooleanOption(sc.ConfigEntry{Key: "RUNNING_IN_DOCKER", Description: "[internal] Running in docker container (don't set this yourself)", DefaultBool: false})
	logOutputFormat := ch.GetStringOption(sc.ConfigEntry{Key: "LOG_OUTPUT_FORMAT", Description: "Log output format (json, text)", DefaultString: "text"})
	disableHmacSigningEndpoint := ch.GetBooleanOption(sc.ConfigEntry{Key: "DISABLE_HMAC_SIGNING_ENDPOINT", Description: "Disable HMAC Request Signing endpoint", DefaultBool: false})
	disableImageEmbedingEndpoints := ch.GetBooleanOption(sc.ConfigEntry{Key: "DISABLE_IMAGE_EMBEDING_ENDPOINTS", Description: "Disable Image Embeding Endpoints", DefaultBool: false})
	enableOnlyHMACEndpoints := ch.GetBooleanOption(sc.ConfigEntry{Key: "ENABLE_ONLY_HMAC_ENDPOINTS", Description: "Enable only HMAC Request Signing endpoint", DefaultBool: false})
	disableMusicBrainzEndpoints := ch.GetBooleanOption(sc.ConfigEntry{Key: "DISABLE_MUSICBRAINZ_ENDPOINTS", Description: "Disable MusicBrainz Endpoints", DefaultBool: false})
	disableEmbeddedMusicBrainz := ch.GetBooleanOption(sc.ConfigEntry{Key: "DISABLE_MUSICBRAINZ_EMBEDDING", Description: "Disable the enrichment with MusicBrainz data", DefaultBool: false})
	disableSwaggerUI := ch.GetBooleanOption(sc.ConfigEntry{Key: "DISABLE_SWAGGER_UI", Description: "Disable Swagger UI", DefaultBool: false})
	// Version Information
	tawnyVers := ch.GetStringOption(sc.ConfigEntry{Key: "INTERNAL_VERSION", Description: "[internal] TAWNY Version (don't set this yourself)", DefaultString: "unspecified"})
	tawnyRev := ch.GetStringOption(sc.ConfigEntry{Key: "INTERNAL_REVISION", Description: "[internal] TAWNY Revision (don't set this yourself)", DefaultString: "unspecified"})
	disableGinVersionPublication := ch.GetBooleanOption(sc.ConfigEntry{Key: "GIN_VERSION_EXPOSE", Description: "do you want to expose the gin version?", DefaultBool: false})
	dontexposesensitiveinformation := ch.GetBooleanOption(sc.ConfigEntry{Key: "DONT_EXPOSE_SENSITIVE_INFORMATION", Description: "Do you want to hide potentially sensitive information", DefaultBool: false})
	// Operator Information
	operator := ch.GetStringOption(sc.ConfigEntry{Key: "OPERATOR", Description: "Operator name, if you wish to publish it", DefaultString: "unspecified"})
	operatorContact := ch.GetStringOption(sc.ConfigEntry{Key: "OPERATOR_CONTACT", Description: "Operator contact (email, link)", DefaultString: "unspecified"})
	operatorPrivacyPolicy := ch.GetStringOption(sc.ConfigEntry{Key: "OPERATOR_PRIVACY_POLICY", Description: "A link to your privacy policy (may be required in some jurisdictions)", DefaultString: "unspecified"})
	operatorImprint := ch.GetStringOption(sc.ConfigEntry{Key: "OPERATOR_IMPRINT", Description: "A link to your imprint (may be required in some jurisdictions, especially Germany)", DefaultString: "unspecified"})
	// Source Code Information
	sourceCodeURL := ch.GetStringOption(sc.ConfigEntry{Key: "SOURCE_CODE_URL", Description: "Source Code URL, please change if you forked or modified the source", DefaultString: "https://github.com/dozro/tawny"})
	sourceAbuseContact := ch.GetStringOption(sc.ConfigEntry{Key: "SOURCE_ABUSE_CONTACT", Description: "Abuse contact, please change if you forked or modified the source", DefaultString: "abuse@itsrye.dev"})
	// Caching
	//redisEnable := ch.GetBooleanOption(sc.ConfigEntry{Key: "ENABLE_REDIS", DefaultBool: false, Description: "Enable Redis"})
	disableCaching := ch.GetBooleanOption(sc.ConfigEntry{Key: "DISABLE_CACHING", DefaultBool: false, Description: "Disable Caching"})
	localCachePath := ch.GetStringOption(sc.ConfigEntry{Key: "LOCAL_CACHE_PATH", DefaultString: "/tmp/tawny/cache", Description: "Path to cache to"})
	saveCacheToFs := ch.GetBooleanOption(sc.ConfigEntry{Key: "SAVE_CACHE_TO_FS", DefaultBool: true, Description: "save cache to fs"})
	ch.ParseFlags()
	return &ServerConfig{
		ApiPort:      *apiport,
		ApiHost:      *apihost,
		ApiBasePath:  *apibasepath,
		HmacSecret:   *hmacsecret,
		LastFMAPIKey: *lastfmapikey,
		DebugMode:    *debugmode,
		DevelopMode:  *developmode,
		ReleaseMode:  *releasemode,
		CachingConfig: CachingConfig{
			EnableRedis:    false,
			DisableCaching: *disableCaching,
			LocalCachePath: *localCachePath,
			SaveCacheToFS:  *saveCacheToFs,
		},
		DisabledEndpoints: ServerDisabledEndpoints{
			DisableHMACSigningEndpoint:    *disableHmacSigningEndpoint,
			DisableMusicBrainzEndpoints:   *disableMusicBrainzEndpoints,
			DisableImageEmbeddedEndpoints: *disableImageEmbedingEndpoints,
			EnableOnlyHMACEndpoints:       *enableOnlyHMACEndpoints,
			DisableSwaggerUI:              *disableSwaggerUI,
		},
		ExtendedServerConfig: ExtendedServerConfig{
			RunningInDocker:            *runningInDocker,
			LogOutputFormat:            *logOutputFormat,
			DisableEmbeddedMusicBrainz: *disableEmbeddedMusicBrainz,
			TawnyVersion:               *tawnyVers,
			TawnyRevision:              *tawnyRev,
			DisableGinVersionPublished: *disableGinVersionPublication,
			HideSensitiveInformation:   *dontexposesensitiveinformation,
			SourceCodeURL:              *sourceCodeURL,
			SourceAbuseContact:         *sourceAbuseContact,
		},
		ServerOperator: ServerOperatorInfo{
			OperatorName:     *operator,
			OperatorContact:  *operatorContact,
			PrivacyPolicyURL: *operatorPrivacyPolicy,
			ImprintURL:       *operatorImprint,
		},
	}
}
