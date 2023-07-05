package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAllStories(t *testing.T) {
	t.Run("should return correct number of stories", func(t *testing.T) {
		stories := GetAllStories()

		assert.Len(t, stories, 2, "Expected number of stories to be 2")
	})
}

func TestCreateStory(t *testing.T) {
	t.Run("should create a new story", func(t *testing.T) {
		story := Story{
			StoryID:      3,
			UserID:       1,
			StoryContent: "Story 3 Content",
			CreatedAt:    time.Now(),
			DeletedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		CreateStory(story)

		stories := GetAllStories()

		assert.Len(t, stories, 3, "Expected number of stories to be 3")

		lastStory := stories[len(stories)-1]
		assert.Equal(t, 3, lastStory.StoryID, "Expected story ID to be 3")
		assert.Equal(t, 1, lastStory.UserID, "Expected user ID to be 1")
		assert.Equal(t, "Story 3 Content", lastStory.StoryContent, "Expected story content to be 'Story 3 Content'")
		// You can add assertions for CreatedAt, DeletedAt, and UpdatedAt if necessary
	})
}
