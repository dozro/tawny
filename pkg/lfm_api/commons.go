package lfm_api

import (
	"fmt"
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
