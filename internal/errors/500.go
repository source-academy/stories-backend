package apierrors

import (
	"net/http"
)

type InternalServerError struct {
	Message string
}

func (e InternalServerError) Error() string {
	return e.Message
}

func (e InternalServerError) HTTPStatusCode() int {
	return http.StatusBadRequest
}
