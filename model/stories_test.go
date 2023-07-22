package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// FIXME: Coupling with the other operations in the users database
// func TestGetAllStories(t *testing.T) {
// 	t.Run("should return correct initial number of stories", func(t *testing.T) {
// 		db, cleanUp := setupDBConnection(t, dbConfig)
// 		defer cleanUp(t)

// 		db.Exec("DELETE FROM stories")
// 		stories := GetAllStories(db)
// 		assert.Len(t, stories, 0, "Expected initial number of stories to be 0")

// 		story := Story{
// 			AuthorID: 1,
// 			Content:  "# Hi\n\nThis is a test story.",
// 		}
// 		CreateStory(db, &story)
// 		stories = GetAllStories(db)
// 		assert.Len(t, stories, 1, "Expected number of stories to be 1 after adding 1 story")
// 	})
// }

func TestCreateStory(t *testing.T) {
	t.Run("should increase the total story count", func(t *testing.T) {
		db, cleanUp := setupDBConnection(t, dbConfig)
		defer cleanUp(t)

		initialStories := GetAllStories(db)

		// We need to first create a user due to the foreign key constraint
		user := User{
			GithubUsername: "testUsername",
			GithubID:       123,
		}
		_ = CreateUser(db, &user)

		story := Story{
			AuthorID: user.ID,
			Content:  "# Hi\n\nThis is a test story.",
		}
		CreateStory(db, &story)

		newStories := GetAllStories(db)
		assert.Len(t, newStories, len(initialStories)+1, "Expected number of stories to increase by 1")

		var lastStory Story
		db.Model(&Story{}).Last(&lastStory)

		assert.Equal(t, story.ID, lastStory.ID, "Expected the story ID to be updated")
		assert.Equal(t, story.AuthorID, lastStory.AuthorID, fmt.Sprintf(expectCreateEqualMessage, "story"))
		assert.Equal(t, story.Content, lastStory.Content, fmt.Sprintf(expectCreateEqualMessage, "story"))
	})
}

func TestGetStoryByID(t *testing.T) {
	t.Run("should get the correct story", func(t *testing.T) {
		db, cleanUp := setupDBConnection(t, dbConfig)
		defer cleanUp(t)

		stories := []*Story{
			{AuthorID: 1, Content: "The quick"},
			{AuthorID: 1, Content: "brown fox"},
			{AuthorID: 1, Content: "jumps over"},
		}

		for _, storyToAdd := range stories {
			CreateStory(db, storyToAdd)
		}

		for _, story := range stories {
			// FIXME: Don't use typecast
			dbStory := GetStoryByID(db, int(story.ID))
			assert.Equal(t, story.ID, dbStory.ID, fmt.Sprintf(expectReadEqualMessage, "story"))
			assert.Equal(t, story.AuthorID, dbStory.AuthorID, fmt.Sprintf(expectReadEqualMessage, "story"))
			assert.Equal(t, story.Content, dbStory.Content, fmt.Sprintf(expectReadEqualMessage, "story"))
		}
	})
}
