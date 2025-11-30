package apiError

type ApiError struct {
	HttpCode          int          `json:"http_code"`
	InternalErrorCode ApiErrorCode `json:"internal_error_code"`
	InternalErrorMsg  string       `json:"internal_error_msg"`
	Message           string       `json:"message"`
	Data              interface{}  `json:"data"`
	Success           bool         `json:"success"`
	Date              string       `json:"date"`
}
