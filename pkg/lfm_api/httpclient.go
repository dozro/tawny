package lfm_api

import "net/http"

const userAgent = "Project Tawny (github.com/dozro/tawny)"

var httpClient = &http.Client{}

func doHttpGetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/xml")
	return httpClient.Do(req)
}
