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
	MusicBrainzApiError
	MusicBrainzLookupDisabledByConfig
	InternalTawnyError
	SelectedImageEncodingNotSupported
	ArrayIsUnexpectedEmpty
)

var ApiErrorCodeNames = map[ApiErrorCode]string{
	NoError:                           "No Error",
	InvalidJson:                       "Invalid JSON",
	InvalidHMAC:                       "Invalid HMAC",
	MissingApiKeyInRequest:            "Missing API key in Request",
	JsonNotMatchingExpectedSchema:     "JSON not matching expected schema",
	InvalidBody:                       "Invalid Body",
	EndpointDisabledByConfig:          "Endpoint disabled by config",
	MusicBrainzApiError:               "MusicBrainz API error",
	MusicBrainzLookupDisabledByConfig: "MusicBrainz lookup disabled by config",
	InternalTawnyError:                "Internal tawny error",
	SelectedImageEncodingNotSupported: "Selected image encoding not supported",
	ArrayIsUnexpectedEmpty:            "Array is unexpected empty",
}

func (aec ApiErrorCode) String() string {
	return ApiErrorCodeNames[aec]
}
