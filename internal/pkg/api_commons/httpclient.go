package api_commons

import (
	"fmt"
	"net/http"
)

var userAgent = "Tawny/0.0.3 (Linux; +https://github.com/dozro/tawny; +abuse@itsrye.dev)"

var httpClient = &http.Client{}

func doHttpGetRequestJSONWithAuth(url, auth string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", auth)
	req.Header.Set("Accept", "application/json")
	return httpClient.Do(req)
}

func doHttpGetRequestXMLWithAuth(url, auth string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", auth)
	req.Header.Set("Accept", "application/xml")
	return httpClient.Do(req)
}

type UserAgentSetupArgs struct {
	Version            string
	Repository         string
	SourceAbuseContact string
	OperatorContact    string
	OperatorName       string
	OperatorImprint    string
}

func SetUpUserAgent(args UserAgentSetupArgs) {
	userAgent = fmt.Sprintf(
		"Tawny/%s (+%s; abuse: %s; operator: %s; contact: %s; imprint: %s)",
		args.Version,
		args.Repository,
		args.SourceAbuseContact,
		args.OperatorName,
		args.OperatorContact,
		args.OperatorImprint,
	)
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
