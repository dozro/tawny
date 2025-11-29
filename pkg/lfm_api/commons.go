package lfm_api

import (
	"fmt"
)

func pageLimitAK(baseUrl, method, username, apiKey string, limit, page int) string {
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

func fromToAK(baseUrl, method string, args FromToAKArgs) string {
	if args.From == -1 && args.To == -1 {
		return fmt.Sprintf("%s?method=%s&user=%s", baseUrl, method, args.UserName)
	} else if args.From != -1 && args.To == -1 {
		return fmt.Sprintf("%s?method=%s&user=%s&from=%d", baseUrl, method, args.UserName, args.From)
	} else if args.From == -1 && args.To != -1 {
		return fmt.Sprintf("%s?method=%s&user=%s&to=%d", baseUrl, method, args.UserName, args.To)
	} else {
		return fmt.Sprintf("%s?method=%s&user=%s&to=%d&from=%d", baseUrl, method, args.UserName, args.To, args.From)
	}
}

type FromToAKArgs struct {
	UserName string `xml:"user"`
	From     int    `xml:"from"`
	To       int    `xml:"to"`
	ApiKey   string `xml:"api_key"`
}
