package http_rego

import "net/http"

type APIError struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (a *APIError) Error() string {
	return a.Message
}

var (
	ErrResourceNotFound = APIError{
		Status:  http.StatusNotFound,
		Code:    http.StatusText(http.StatusNotFound),
		Message: "resource not found",
	}
	ErrMalformedBody = APIError{
		Status:  http.StatusBadRequest,
		Code:    http.StatusText(http.StatusBadRequest),
		Message: "request body is not valid",
	}
	ErrKeyExists = APIError{
		Status:  http.StatusBadRequest,
		Code:    http.StatusText(http.StatusBadRequest),
		Message: "key that wans to be stored has already exists",
	}
)
