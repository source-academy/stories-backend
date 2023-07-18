package apierrors

import (
	"net/http"
)

type ClientNotFoundError struct {
	message string
}

func (e *ClientNotFoundError) Error() string {
	return e.message
}

func (e *ClientNotFoundError) HTTPStatusCode() int {
	return http.StatusNotFound
}
