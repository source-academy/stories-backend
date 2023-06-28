package model

// can use :  `go test -v ./model` to run test

import (
	"testing"

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
			ID:     3,
			Title:  "Story 3",
			Author: "Username3",
		}

		CreateStory(story)

		stories := GetAllStories()

		assert.Len(t, stories, 3, "Expected number of stories to be 3")

		lastStory := stories[len(stories)-1]
		assert.Equal(t, 3, lastStory.ID, "Expected story ID to be 3")
		assert.Equal(t, "Story 3", lastStory.Title, "Expected story title to be 'Story 3'")
		assert.Equal(t, "Username3", lastStory.Author, "Expected story author to be 'Username3'")
	})
}
