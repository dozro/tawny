package api_commons

import (
	"fmt"
	"net/http"
)

var userAgent = "Tawny/0.0.3 (Linux; +https://github.com/dozro/tawny; +abuse@itsrye.dev)"

const stringsUa = "User-Agent"
const stringsAuth = "Authorization"
const stringsAccepts = "Accept"
const stringsAcceptsXML = "application/xml"
const stringsAcceptsJSON = "application/json"

var httpClient = &http.Client{}

func doHttpGetRequestJSONWithAuth(url, auth string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(stringsUa, userAgent)
	req.Header.Set(stringsAuth, auth)
	req.Header.Set(stringsAccepts, stringsAcceptsJSON)
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
	req.Header.Set(stringsUa, userAgent)
	req.Header.Set(stringsAccepts, stringsAcceptsXML)
	return httpClient.Do(req)
}

func SetHttpClient(client *http.Client) {
	httpClient = client
}
