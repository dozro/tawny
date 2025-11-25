package embed

import (
	"net/http"
	"time"
)

const userAgent = "Tawny-Image-Embed/0.0.1 (linux;github.com/dozro/tawny;+abuse@itsrye.dev)"

var httpClient = &http.Client{}

func doHttpGetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "image/png")
	return httpClient.Do(req)
}

func init() {
	httpClient = http.DefaultClient
	httpClient.Timeout = time.Second * 30
}
