package lfm_api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"lastfm-proxy/pkg/lfm_types"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (User) GetRecentTracks(args UserGetLovedTracksArgs) (*lfm_types.UserGetRecentTracks, error) {
	var apiUrl string
	if -1 != args.Limit && -1 != args.Page {
		apiUrl = fmt.Sprintf("%s?method=user.getrecenttracks&user=%s&api_key=%s&limit=%d&page=%d", baseUrl, args.UserName, args.ApiKey, args.Limit, args.Page)
	} else if -1 != args.Limit {
		apiUrl = fmt.Sprintf("%s?method=user.getrecenttracks&user=%s&api_key=%s&limit=%d", baseUrl, args.UserName, args.ApiKey, args.Limit)
	} else if -1 != args.Page {
		apiUrl = fmt.Sprintf("%s?method=user.getrecenttracks&user=%s&api_key=%s&page=%d", baseUrl, args.UserName, args.ApiKey, args.Page)
	} else {
		apiUrl = fmt.Sprintf("%s?method=user.getrecenttracks&user=%s&api_key=%s", baseUrl, args.UserName, args.ApiKey)
	}
	log.Debugf("apiUrl: %s", apiUrl)
	resp, err := doHttpGetRequest(apiUrl)
	log.Debugf("Response from API: %v", resp)
	if err != nil {
		log.Errorf("Error getting user info: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data lfm_types.WrappedUserGetRecentTracks
	err = xml.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data.RecentTracks, nil
}
