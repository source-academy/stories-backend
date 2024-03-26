package storyparams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	negativePinOrder := -1
	tests := []struct {
		name    string
		params  Create
		wantErr bool
	}{
		{"valid input", Create{AuthorID: 1, Title: "Test Title", Content: "Test Content"}, false},
		{"missing authorId", Create{Title: "Test Title", Content: "Test Content"}, true},
		{"empty title", Create{AuthorID: 1, Content: "Test Content"}, true},
		{"empty content", Create{AuthorID: 1, Title: "Test Title"}, true},
		{"negative pinOrder", Create{AuthorID: 1, Title: "Test Title", Content: "Test Content", PinOrder: &negativePinOrder}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
