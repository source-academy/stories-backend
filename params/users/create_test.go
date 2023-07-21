package userparams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("should do nothing for now", func(t *testing.T) {})
}

func TestToModel(t *testing.T) {
	t.Run("should create a user model with the correct values", func(t *testing.T) {
		params := Create{
			Username:      "testUsername",
			LoginProvider: "github",
		}
		model := params.ToModel()
		assert.Equal(t, params.Username, model.Username)
		assert.Equal(t, params.LoginProvider, model.LoginProvider.String())
	})
	// TODO: Add test for invalid login provider
}
