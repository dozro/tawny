package server

import (
	"encoding/json"
)

type HmacSignedRequest struct {
	Signature string `json:"signature"`
	// Raw JSON to avoid issues with encoding while verifying the request
	Request json.RawMessage `json:"request"`
}

type HmacBase64SignedRequest struct {
	Signature string `json:"signature"`
	Request   []byte `json:"request"`
}

func HmacSignedRequestToBase64(request HmacSignedRequest) HmacBase64SignedRequest {
	signedRequest := HmacBase64SignedRequest{
		Signature: request.Signature,
		Request:   request.Request,
	}
	return signedRequest
}

func Base64ToHmacSignedRequest(b64 HmacBase64SignedRequest) (HmacSignedRequest, error) {
	signedReq := HmacSignedRequest{
		Signature: b64.Signature,
		Request:   json.RawMessage(b64.Request),
	}
	return signedReq, nil
}

type HmacProxyRequest struct {
	Method        string                        `json:"method"`
	ApiIdentifier string                        `json:"api_identifier"`
	ApiParameters HmacProxyRequestApiParameters `json:"api_parameters"`
	Body          string                        `json:"body"`
}

type HmacProxyRequestApiParameters struct {
	Username string `json:"username"`
}

type HmacProxyResponseWrapper struct {
	HttpCode int               `json:"http_code"`
	Response HmacProxyResponse `json:"response"`
}

type HmacProxyResponse struct {
	Schema string          `json:"$schema"`
	Data   json.RawMessage `json:"data"`
}
