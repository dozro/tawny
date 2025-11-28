package api_commons

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dozro/tawny/internal/pkg/security"
	log "github.com/sirupsen/logrus"
)

func FetchJSON[T any](url string, authHeader string) (T, error) {
	log.Debugf("fetching json from %s with auth header: %s", url, security.MaskAPIKey(authHeader))
	var zero T

	resp, err := doHttpGetRequestJSONWithAuth(url, authHeader)
	if err != nil {
		log.Errorf("http request failed: %v; url was: %s", err.Error(), url)
		return zero, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("unexpected status: %s; url was: %s", resp.Status, url)
		return zero, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)

	var result T
	if err := decoder.Decode(&result); err != nil {
		return zero, fmt.Errorf("json decode failed: %w", err)
	}

	return result, nil
}
