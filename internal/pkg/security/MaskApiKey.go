package security

import (
	"math/rand"
	"net/url"
	"strings"
)

func MaskAPIKey(key string) string {
	randnum := rand.Intn(20-5) + 5
	if len(key) <= 8 {
		return strings.Repeat("X", randnum)
	}

	return key[:4] + strings.Repeat("X", randnum)
}

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
	for i, part := range parts {
		if len(part) > 10 && (strings.HasPrefix(part, "sk-") || strings.HasPrefix(part, "ak-")) {
			parts[i] = MaskAPIKey(part)
		}
	}
	u.Path = strings.Join(parts, "/")

	return u.String()
}
