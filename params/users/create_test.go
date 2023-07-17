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
			GithubUsername: "testUsername",
			GithubID:       123,
		}
		model := params.ToModel()
		assert.Equal(t, params.GithubUsername, model.GithubUsername)
		assert.Equal(t, params.GithubID, model.GithubID)
	})
}
