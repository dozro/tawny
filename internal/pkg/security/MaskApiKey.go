package security

import (
	"math/rand"
	"net/url"
	"strings"
)
	"net/url"
	"strings"
)

const maskLength = 12

func MaskAPIKey(key string) string {
	if len(key) <= 8 {
		return strings.Repeat("X", maskLength)
	}

	return key[:4] + strings.Repeat("X", maskLength)
}

// MaskURLKey masks sensitive API keys within URLs for safe logging.
// It masks query parameters named "api_key" and path segments that
// look like API keys (prefixed with "sk-" or "ak-" and longer than 10 chars).
func MaskURLKey(fullURL string) string {
	u, err := url.Parse(fullURL)
	if err != nil {
		return strings.Repeat("*", min(20, len(fullURL)))
	}

	q := u.Query()
	if keyVal := q.Get("api_key"); keyVal != "" {
		q.Set("api_key", MaskAPIKey(keyVal))
		u.RawQuery = q.Encode()
	}

	parts := strings.Split(u.Path, "/")
	// Mask path segments that match common API key patterns.
	const minAPIKeyLength = 8 // Chosen to catch most API keys, configurable as needed.
	for i, part := range parts {
		// Check for common API key prefixes and patterns.
		if len(part) > minAPIKeyLength &&
			(strings.HasPrefix(part, "sk-") ||
				strings.HasPrefix(part, "ak-") ||
				strings.HasPrefix(part, "api_key") ||
				strings.HasPrefix(part, "apikey") ||
				strings.HasPrefix(part, "token") ||
				strings.Contains(part, "apikey") ||
				strings.Contains(part, "api_key")) {
			parts[i] = MaskAPIKey(part)
		}
	}
	u.Path = strings.Join(parts, "/")

	return u.String()
}
