package tawny_sdk

import (
	"net/url"
)

type TawnyCreationArgs struct {
	LastFMApiKey  string  `json:"lastFMApiKey"`
	TawnyEndPoint url.URL `json:"tawnyEndPoint"`
}

type Tawny struct {
	LastFMApiKey  string  `json:"lastFMApiKey"`
	TawnyEndPoint url.URL `json:"tawnyEndPoint"`
}

func (t Tawny) getApiBaseUrlString() string {
	return t.TawnyEndPoint.String()
}

func (Tawny) NewTawny(args TawnyCreationArgs) *Tawny {
	return &Tawny{
		LastFMApiKey:  args.LastFMApiKey,
		TawnyEndPoint: args.TawnyEndPoint,
	}
}
