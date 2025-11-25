package proxy_config

import "os"

type ProxyConfig struct {
	HmacSecret   string
	LastFMAPIKey string
}

func GetProxyConfig() *ProxyConfig {
	hmacSecret := os.Getenv("TAWNY_HMAC_SECRET")
	lastfmApiKey := os.Getenv("TAWNY_LASTFM_API_KEY")
	return &ProxyConfig{
		HmacSecret:   hmacSecret,
		LastFMAPIKey: lastfmApiKey,
	}
}
