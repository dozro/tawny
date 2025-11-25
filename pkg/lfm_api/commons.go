package lfm_api

import (
	"encoding/xml"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func pageLimitAK(baseUrl string, method string, username string, apiKey string, limit int, page int) string {
	if -1 != limit && -1 != page {
		return fmt.Sprintf("%s?method=%s&user=%s&api_key=%s&limit=%d&page=%d", baseUrl, method, username, apiKey, limit, page)
	} else if -1 != limit {
		return fmt.Sprintf("%s?method=%s&user=%s&api_key=%s&limit=%d", baseUrl, method, username, apiKey, limit)
	} else if -1 != page {
		return fmt.Sprintf("%s?method=%s&user=%s&api_key=%s&page=%d", baseUrl, method, username, apiKey, page)
	} else {
		return fmt.Sprintf("%s?method=%s&user=%s&api_key=%s", baseUrl, method, username, apiKey)
	}
}

func fetchXML[T any](url string) (T, error) {
	log.Debug("fetchXML")
	var zero T

	resp, err := doHttpGetRequest(url)
	if err != nil {
		log.Errorf("http request failed: %v; url was: %s", err.Error(), url)
		return zero, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("unexpected status: %s; url was: %s", resp.Status, url)
		return zero, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	decoder := xml.NewDecoder(resp.Body)

	var result T
	if err := decoder.Decode(&result); err != nil {
		return zero, fmt.Errorf("xml decode failed: %w", err)
	}

	return result, nil
}
