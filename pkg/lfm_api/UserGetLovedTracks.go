package lfm_api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"lastfm-proxy/pkg/lfm_types"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (User) GetLovedTracks(apiKey string, userName string, limit int, page int) (*lfm_types.UserGetLovedTracks, error) {
	var apiUrl string
	if -1 != limit && -1 != page {
		apiUrl = fmt.Sprintf("%s?method=user.getlovedtracks&user=%s&api_key=%s&limit=%d&page=%d", baseUrl, userName, apiKey, limit, page)
	} else if -1 != limit {
		apiUrl = fmt.Sprintf("%s?method=user.getlovedtracks&user=%s&api_key=%s&limit=%d", baseUrl, userName, apiKey, limit)
	} else if -1 != page {
		apiUrl = fmt.Sprintf("%s?method=user.getlovedtracks&user=%s&api_key=%s&page=%d", baseUrl, userName, apiKey, page)
	} else {
		apiUrl = fmt.Sprintf("%s?method=user.getlovedtracks&user=%s&api_key=%s", baseUrl, userName, apiKey)
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

	var data lfm_types.WrappedUserGetLovedTracks
	err = xml.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data.LovedTracks, nil
}
