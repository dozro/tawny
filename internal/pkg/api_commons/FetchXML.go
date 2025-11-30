package api_commons

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/dozro/tawny/internal/pkg/security"
	log "github.com/sirupsen/logrus"
)

func FetchXML[T any](url string) (T, error) {
	log.Debug("fetchXML")
	var zero T

	resp, err := doHttpGetRequest(url)
	if err != nil {
		log.Errorf("http request failed: %v; url was: %s", err.Error(), security.MaskURLKey(url))
		return zero, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("unexpected status: %s; url was: %s", resp.Status, security.MaskURLKey(url))
		return zero, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	decoder := xml.NewDecoder(resp.Body)

	var result T
	if decoder.Decode(&result) != nil {
		return zero, fmt.Errorf("xml decode failed: %w", err)
	}

	return result, nil
}
