package lfm_api

import (
	"net/http"
	"time"
)

const baseUrl = "http://ws.audioscrobbler.com/2.0/"

func init() {
	httpClient = http.DefaultClient
	httpClient.Timeout = time.Second * 30
}
