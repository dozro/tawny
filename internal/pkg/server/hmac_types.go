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
