package api_commons

import "net/http"

const userAgent = "Tawny/0.0.1 (linux;github.com/dozro/tawny;+abuse@itsrye.dev)"

var httpClient = &http.Client{}

func doHttpGetRequestJSONWithAuth(url string, auth string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", auth)
	req.Header.Set("Accept", "application/json")
	return httpClient.Do(req)
}

func doHttpGetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/xml")
	return httpClient.Do(req)
}

func SetHttpClient(client *http.Client) {
	httpClient = client
}
