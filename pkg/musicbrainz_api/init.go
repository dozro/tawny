package musicbrainz_api

import (
	"net/http"
	"time"

	"github.com/dozro/tawny/internal/pkg/api_commons"
)

const baseUrl = "https://musicbrainz.org/ws/2"

func init() {
	httpClient := http.DefaultClient
	httpClient.Timeout = time.Second * 30
	api_commons.SetHttpClient(httpClient)
}
