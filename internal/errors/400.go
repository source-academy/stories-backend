package apierrors

import (
	"net/http"
)

type ClientBadRequestError struct {
	Message string
}

func (e ClientBadRequestError) Error() string {
	return e.Message
}

func (e ClientBadRequestError) HTTPStatusCode() int {
	return http.StatusBadRequest
}
