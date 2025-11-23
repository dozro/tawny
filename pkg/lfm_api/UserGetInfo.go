package lfm_api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"lastfm-proxy/pkg/lfm_types"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (User) GetInfo(apiKey string, userName string) (*lfm_types.UserGetInfo, error) {
	apiUrl := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.getinfo&user=%s&api_key=%s", userName, apiKey)
	log.Debugf("apiUrl: %s", apiUrl)
	resp, err := http.Get(apiUrl)
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

	var data lfm_types.WrappedUserGetInfo
	err = xml.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data.User, nil
}
