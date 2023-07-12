package model

import (
	"testing"
	// "time"

	"github.com/stretchr/testify/assert"
)

func TestGetAllStories(t *testing.T) {
	t.Run("should return all stories", func(t *testing.T) {
		// Retrieve the stories
		retrievedStories := GetAllStories()

		// Assert the number of stories
		assert.Len(t, &retrievedStories, 0, "Expected number of stories to be 0")
	})
}

func TestCreateStory(t *testing.T) {
	t.Run("should create a new story", func(t *testing.T) {

		// Create a new story
		story := Story{
			StoryID:      3,
			AuthorID:     1,
			StoryContent: "Story 3 Content",
			// CreatedAt:    time.Now(),
			// DeletedAt:    time.Now(),
			// UpdatedAt:    time.Now(),
		}

		// Create the story
		CreateStory(story)

		// Retrieve all stories
		stories := GetAllStories()

		// Assert the number of stories
		assert.Len(t, stories, 1, "Expected number of stories to be 1")

		// Assert the properties of the created story
		createdStory := stories[0]
		assert.Equal(t, 3, createdStory.StoryID, "Expected story ID to be 3")
		assert.Equal(t, 1, createdStory.AuthorID, "Expected user ID to be 1")
		assert.Equal(t, "Story 3 Content", createdStory.StoryContent, "Expected story content to be 'Story 3 Content'")
		// You can add assertions for CreatedAt, DeletedAt, and UpdatedAt if necessary
	})
}

func TestGetStoryByID(t *testing.T) {
	t.Run("should retrieve a story by ID", func(t *testing.T) {

		// Retrieve a story by ID
		story := GetStoryByID(3)

		// Assert the retrieved story
		assert.NotNil(t, story, "Expected story to be retrieved")
		assert.Equal(t, 3, story.StoryID, "Expected story ID to be 3")
		assert.Equal(t, 1, story.AuthorID, "Expected user ID to be 1")
		assert.Equal(t, "Story 3 Content", story.StoryContent, "Expected story content to be 'Story 3 Content'")
		// You can add assertions for CreatedAt, DeletedAt, and UpdatedAt if necessary
	})
}

// package model

// import (
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// )

// func TestGetAllStories(t *testing.T) {
// 	t.Run("should return correct number of stories", func(t *testing.T) {
// 		stories := GetAllStories()

// 		assert.Len(t, stories, 2, "Expected number of stories to be 2")
// 	})
// }

// func TestCreateStory(t *testing.T) {
// 	t.Run("should create a new story", func(t *testing.T) {
// 		story := Story{
// 			StoryID:      3,
// 			UserID:       1,
// 			StoryContent: "Story 3 Content",
// 			CreatedAt:    time.Now(),
// 			DeletedAt:    time.Now(),
// 			UpdatedAt:    time.Now(),
// 		}

// 		CreateStory(story)

// 		stories := GetAllStories()

// 		assert.Len(t, stories, 3, "Expected number of stories to be 3")

// 		lastStory := stories[len(stories)-1]
// 		assert.Equal(t, 3, lastStory.StoryID, "Expected story ID to be 3")
// 		assert.Equal(t, 1, lastStory.UserID, "Expected user ID to be 1")
// 		assert.Equal(t, "Story 3 Content", lastStory.StoryContent, "Expected story content to be 'Story 3 Content'")
// 		// You can add assertions for CreatedAt, DeletedAt, and UpdatedAt if necessary
// 	})
// }
