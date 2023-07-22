package userparams

import (
	"testing"

	userenums "github.com/source-academy/stories-backend/internal/enums/users"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("should ensure username is not empty", func(t *testing.T) {
		params := Create{
			Username: "",
		}

		err := params.Validate()
		assert.NotNil(t, err)
	})
	t.Run("should ensure supported and enabled login provider passes", func(t *testing.T) {
		params := Create{
			Username:      "testUsername",
			LoginProvider: userenums.LoginProviderNUSNET.ToString(),
		}
		err := params.Validate()
		assert.Nil(t, err)
	})
	t.Run("should ensure unsupported login provider fails", func(t *testing.T) {
		params := Create{
			Username:      "testUsername",
			LoginProvider: "invalidProvider",
		}
		err := params.Validate()
		assert.NotNil(t, err)
	})
	t.Run("should ensure disabled login provider fails", func(t *testing.T) {
		params := Create{
			Username:      "testUsername",
			LoginProvider: userenums.LoginProviderGitHub.ToString(),
		}
		err := params.Validate()
		assert.NotNil(t, err)
	})
}

func TestToModel(t *testing.T) {
	t.Run("should create a user model with the correct values", func(t *testing.T) {
		params := Create{
			Username:      "testUsername",
			LoginProvider: userenums.LoginProviderNUSNET.ToString(),
		}
		model := params.ToModel()
		assert.Equal(t, params.Username, model.Username)
		assert.Equal(t, params.LoginProvider, model.LoginProvider.ToString())
	})
}
