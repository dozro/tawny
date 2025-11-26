package lfm_api

import (
	"net/http"
	"time"

	"github.com/dozro/tawny/pkg/api_commons"
)

const baseUrl = "http://ws.audioscrobbler.com/2.0/"

func init() {
	httpClient = http.DefaultClient
	httpClient.Timeout = time.Second * 30
	api_commons.SetHttpClient(httpClient)
}
