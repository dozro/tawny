package tawny_sdk

import (
	"net/url"

	log "github.com/sirupsen/logrus"
)

type TawnyCreationArgs struct {
	LastFMApiKey  string  `json:"lastFMApiKey"`
	HMACSecretKey string  `json:"hmacSecretKey"`
	TawnyEndPoint url.URL `json:"tawnyEndPoint"`
}

type Tawny struct {
	LastFMApiKey  string  `json:"lastFMApiKey"`
	HMACSecretKey string  `json:"hmacSecretKey"`
	TawnyEndPoint url.URL `json:"tawnyEndPoint"`
}

func (t Tawny) getApiBaseUrlString() string {
	return t.TawnyEndPoint.String()
}

func (Tawny) NewTawny(args TawnyCreationArgs) *Tawny {
	log.Debug("Setting up new Tawny Client")
	log.Debug(args)
	return &Tawny{
		LastFMApiKey:  args.LastFMApiKey,
		TawnyEndPoint: args.TawnyEndPoint,
		HMACSecretKey: args.HMACSecretKey,
	}
}
