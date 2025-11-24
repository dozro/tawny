package lfm_api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dozro/tawny/pkg/lfm_types"

	log "github.com/sirupsen/logrus"
)

func (User) GetRecentTracks(args UserGetArgsWithLimitPage) (*lfm_types.UserGetRecentTracks, error) {
	apiUrl := pageLimitAK(baseUrl, "user.getRecentTracks", args.UserName, args.ApiKey, args.Limit, args.Page)
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

	for i := range data.RecentTracks.Track {
		data.RecentTracks.Track[i].Brainz()
	}

	return &data.RecentTracks, nil
}
