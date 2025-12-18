package musicbrainz_api

import (
	"net/http"
	"regexp"
	"time"

	"github.com/dozro/tawny/internal/pkg/api_commons"
)

const baseUrl = "https://musicbrainz.org/ws/2"

var mbidRegexp *regexp.Regexp

func init() {
	httpClient := http.DefaultClient
	httpClient.Timeout = time.Second * 30
	api_commons.SetHttpClient(httpClient)
	mbidRegexp = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)
}
