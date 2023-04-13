package enum

import "net/http"

const (
	ErrResourceMissing     = "resource_missing"
	ErrInternalServerError = "internal_server_error"
	ErrHeaderInvalid       = "header_invalid"
	ErrRequestBodyInvalid  = "request_body_invalid"
)

func NewErrorCodeMap() map[string]int {
	var errorCodeMap = make(map[string]int)
	errorCodeMap[ErrResourceMissing] = http.StatusNotFound
	errorCodeMap[ErrInternalServerError] = http.StatusInternalServerError
	errorCodeMap[ErrHeaderInvalid] = http.StatusBadRequest
	errorCodeMap[ErrRequestBodyInvalid] = http.StatusBadRequest

	return errorCodeMap
}
