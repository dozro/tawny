package tawny_sdk

import (
	"encoding/base64"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"gitlab.com/rye_tawny/api_commons"
	"gitlab.com/rye_tawny/hmac_types"
	"gitlab.com/rye_tawny/security"
)

func executeHmac[T any](parameters hmac_types.HmacProxyRequestApiParameters, apiId, apiUrl, hmacsec string) (*T, error) {
	request := hmac_types.HmacProxyRequest{
		Method:        "GET",
		ApiIdentifier: apiId,
		ApiParameters: parameters,
	}

	packed, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	urlencReq := base64.StdEncoding.EncodeToString(packed)
	log.Debugf("HMAC Request: %s", string(packed))
	signature := security.GenerateHMAC(hmacsec, string(packed))

	signedReq := hmac_types.HmacBase64SignedRequest{
		Signature: signature,
		Request:   []byte(urlencReq),
	}
	jsonSignedReq, err := json.Marshal(signedReq)
	log.Debugf("signedReq: %s", string(jsonSignedReq))

	// Self Test
	if !security.VerifyHMAC(hmacsec, string(packed), signature) {
		log.Errorf("failed to self generated HMAC")
		log.Fatalf("FAILED SELF-VERIFICATION. THIS IS A BUG. Generated Sig: %s for %v (%s)", signature, request, urlencReq)
	}

	log.Debugf("Signature: %s", signature)

	ret, err := api_commons.FetchJSONWithPostWithoutAuth[T](apiUrl, jsonSignedReq)
	if err != nil {
		log.Error(err)
		log.Infof("Args were sig=%s, req=%s", signature, urlencReq)
		return nil, err
	}
	return &ret, nil
}
