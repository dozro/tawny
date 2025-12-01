package tawny_sdk

import (
	"fmt"

	lfmtypes "github.com/dozro/tawny_lfm_types"
	"gitlab.com/rye_tawny/api_commons"
	"gitlab.com/rye_tawny/hmac_types"
)

func (t Tawny) GetNowListeningFor(username string) (*lfmtypes.LFMTrack, error) {
	apiUrl := fmt.Sprintf("%s/user/%s/tracks/current", t.getApiBaseUrlString(), username)
	ct, err := api_commons.FetchJSON[lfmtypes.UserGetRecentTracks](apiUrl, t.LastFMApiKey)
	if err != nil {
		return nil, err
	}
	return &ct.Track[0], nil
}

func (t Tawny) SecureNowListeningFor(username string) (*lfmtypes.LFMTrack, error) {

	apiUrl := fmt.Sprintf("%s/hmac/execute?is_base64=true", t.getApiBaseUrlString())
	ct, err := executeHmac[lfmtypes.UserGetRecentTracks](hmac_types.HmacProxyRequestApiParameters{Username: username}, "user/tracks/current", apiUrl, t.HMACSecretKey)
	if err != nil {
		return nil, err
	}
	return &ct.Track[0], nil
}
