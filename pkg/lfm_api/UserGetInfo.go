package lfm_api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dozro/tawny/pkg/lfm_types"

	log "github.com/sirupsen/logrus"
)

type UserGetInfoArgs struct {
	ApiKey   string
	UserName string
}

func (User) GetInfo(args UserGetInfoArgs) (*lfm_types.UserGetInfo, error) {
	apiUrl := fmt.Sprintf("%s?method=user.getinfo&user=%s&api_key=%s", baseUrl, args.UserName, args.ApiKey)
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

	var data lfm_types.WrappedUserGetInfo
	err = xml.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data.User, nil
}
