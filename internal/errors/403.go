package apierrors

import (
	"net/http"
)

type ClientForbiddenError struct {
	Message string
}

func (e ClientForbiddenError) Error() string {
	return e.Message
}

func (e ClientForbiddenError) HTTPStatusCode() int {
	return http.StatusForbidden
}
