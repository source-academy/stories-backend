package storyparams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("should do nothing for now", func(t *testing.T) {})
}

func TestToModel(t *testing.T) {
	t.Run("should create a story model with the correct values", func(t *testing.T) {
		params := Create{
			AuthorID: 1,
			Content:  "# Hi\n\nThis is a test story.",
		}
		model := params.ToModel(nil)
		assert.Equal(t, params.AuthorID, model.AuthorID)
		assert.Equal(t, params.Content, model.Content)
	})
}
