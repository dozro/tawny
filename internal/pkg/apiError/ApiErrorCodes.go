package apiError

type ApiErrorCode int

const (
	NoError ApiErrorCode = iota
	InvalidJson
	InvalidHMAC
	MissingApiKeyInRequest
	JsonNotMatchingExpectedSchema
	InvalidBody
	EndpointDisabledByConfig
)

var ApiErrorCodeNames = map[ApiErrorCode]string{
	NoError:                       "NoError",
	InvalidJson:                   "invalid_json",
	InvalidHMAC:                   "invalid_hmac",
	MissingApiKeyInRequest:        "missing_api_key_in_request",
	JsonNotMatchingExpectedSchema: "json_not_matching_expected_schema",
	InvalidBody:                   "invalid_body",
	EndpointDisabledByConfig:      "endpoint_disabled_by_config",
}

func (aec ApiErrorCode) String() string {
	return ApiErrorCodeNames[aec]
}
