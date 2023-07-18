package apierrors

import (
	"net/http"
)

type ClientNotFoundError struct {
	Message string
}

func (e ClientNotFoundError) Error() string {
	return e.Message
}

func (e ClientNotFoundError) HTTPStatusCode() int {
	return http.StatusNotFound
}
