package router

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func handleAPIError(handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {

func TestHandleAPIError(t *testing.T) {
	mockController := func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
	t.Run("Should return a http.HandlerFunc", func(t *testing.T) {
		var handler interface{} = handleAPIError(mockController)
		assert.NotNil(t, handler, "Should not be nil")
		_, ok := handler.(http.HandlerFunc)
		assert.True(t, ok, "Should be a http.HandlerFunc")
	})
}
