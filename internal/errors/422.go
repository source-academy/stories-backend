package apierrors

import (
	"net/http"
)

type ClientUnprocessableEntityError struct {
	Message string
}

func (e ClientUnprocessableEntityError) Error() string {
	return e.Message
}

func (e ClientUnprocessableEntityError) HTTPStatusCode() int {
	return http.StatusUnprocessableEntity
}
