package auth

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAndGetUser(t *testing.T) {
	t.Run("should return error if query string is not valid", func(t *testing.T) {
		qs := "invalid"
		_, err := validateAndGetUser(qs, nil)
		assert.NotNil(t, err)
	})

	t.Run("should return error if query string does not contain required fields", func(t *testing.T) {
		values := url.Values{
			"someKey": []string{"someValue"},
			"another": []string{"anotherValue"},
		}
		qs := values.Encode()
		_, err := validateAndGetUser(qs, nil)
		assert.NotNil(t, err)
	})
	t.Run("should return error if login provider is invalid or unsupported", func(t *testing.T) {
		values := url.Values{
			usernameKey:      []string{"someValue"},
			loginProviderKey: []string{"invalid"},
		}
		qs := values.Encode()
		_, err := validateAndGetUser(qs, nil)
		assert.NotNil(t, err)
	})
	// TODO: DB-related tests
}
