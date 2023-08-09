package model

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	userenums "github.com/source-academy/stories-backend/internal/enums/users"
	"github.com/source-academy/stories-backend/internal/testutils"
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

var (
	errInvalidForeignKey = pgconn.PgError{
		Code: "23503",
	}
	errNonNullable = pgconn.PgError{
		Code: "23502",
	}
)

func TestCreateStory(t *testing.T) {
	t.Run("should increase the total story count", func(t *testing.T) {
		db, cleanUp := testutils.SetupDBConnection(t, dbConfig, migrationPath)
		defer cleanUp(t)

		initialStories, err := GetAllStories(db)
		assert.Nil(t, err, "Expected no error when getting all stories")

		// We need to first create a user and a group due to the foreign key constraint
		user := User{
			Username:      "testStoryAuthor",
			LoginProvider: userenums.LoginProvider(rand.Int31()),
		}
		_ = CreateUser(db, &user)

		group := Group{
			Name: "testGroup",
		}
		_ = CreateGroup(db, &group)

		story := Story{
			AuthorID: user.ID,
			Group:    group,
			Content:  "# Hi\n\nThis is a test story.",
		}
		err = CreateStory(db, &story)
		assert.Nil(t, err, "Expected no error when creating story")

		newStories, err := GetAllStories(db)
		assert.Nil(t, err, "Expected no error when getting all stories")
		assert.Len(t, newStories, len(initialStories)+1, "Expected number of stories to increase by 1")

		var lastStory Story
		db.Model(&Story{}).Last(&lastStory)

		assert.Equal(t, story.ID, lastStory.ID, "Expected the story ID to be updated")
		assert.Equal(t, story.AuthorID, lastStory.AuthorID, fmt.Sprintf(expectCreateEqualMessage, "story"))
		assert.Equal(t, story.GroupID, lastStory.GroupID, fmt.Sprintf(expectCreateEqualMessage, "story"))
		assert.Equal(t, story.Content, lastStory.Content, fmt.Sprintf(expectCreateEqualMessage, "story"))
	})

	t.Run("can create without group", func(t *testing.T) {
		db, cleanUp := testutils.SetupDBConnection(t, dbConfig, migrationPath)
		defer cleanUp(t)

		initialStories, err := GetAllStories(db)
		assert.Nil(t, err, "Expected no error when getting all stories")

		// We need to first create a user and a group due to the foreign key constraint
		user := User{
			Username:      "testStoryAuthor",
			LoginProvider: userenums.LoginProvider(rand.Int31()),
		}
		_ = CreateUser(db, &user)

		story := Story{
			AuthorID: user.ID,
			Content:  "# Hi\n\nThis is a test story.",
		}
		err = CreateStory(db, &story)
		assert.Nil(t, err, "Expected no error when creating story")

		newStories, err := GetAllStories(db)
		assert.Nil(t, err, "Expected no error when getting all stories")
		assert.Len(t, newStories, len(initialStories)+1, "Expected number of stories to increase by 1")

		var lastStory Story
		db.Model(&Story{}).Last(&lastStory)

		assert.Equal(t, story.ID, lastStory.ID, "Expected the story ID to be updated")
		assert.Equal(t, story.AuthorID, lastStory.AuthorID, fmt.Sprintf(expectCreateEqualMessage, "story"))
		assert.Equal(t, story.GroupID, lastStory.GroupID, fmt.Sprintf(expectCreateEqualMessage, "story"))
		assert.Equal(t, story.Content, lastStory.Content, fmt.Sprintf(expectCreateEqualMessage, "story"))
	})

	t.Run("cannot create without author in model", func(t *testing.T) {
		db, cleanUp := testutils.SetupDBConnection(t, dbConfig, migrationPath)
		defer cleanUp(t)

		story := Story{
			Content: "# Hi\n\nThis is a test story.",
		}
		err := CreateStory(db, &story)

		var pgerr *pgconn.PgError
		if assert.ErrorAs(t, err, &pgerr, "Expected error when creating story without Author ID") {
			assert.Equal(t, errInvalidForeignKey.Code, pgerr.Code)
		}
	})
}

func TestGetStoryByID(t *testing.T) {
	t.Run("should get the correct story", func(t *testing.T) {
		db, cleanUp := testutils.SetupDBConnection(t, dbConfig, migrationPath)
		defer cleanUp(t)

		// We need to first create a user due to the foreign key constraint
		user := User{
			Username:      "testMultipleStoriesAuthor",
			LoginProvider: userenums.LoginProvider(rand.Int31()),
		}
		_ = CreateUser(db, &user)

		group := Group{
			Name: "testGroup",
		}
		_ = CreateGroup(db, &group)

		stories := []*Story{
			{AuthorID: user.ID, Group: group, Content: "The quick"},
			{AuthorID: user.ID, Group: group, Content: "brown fox"},
			{AuthorID: user.ID, Group: group, Content: "jumps over"},
		}

		for _, storyToAdd := range stories {
			_ = CreateStory(db, storyToAdd)
		}

		for _, story := range stories {
			// FIXME: Don't use typecast
			dbStory, err := GetStoryByID(db, int(story.ID))
			assert.Nil(t, err, "Expected no error when getting story with valid ID")
			assert.Equal(t, story.ID, dbStory.ID, fmt.Sprintf(expectReadEqualMessage, "story"))
			assert.Equal(t, story.AuthorID, dbStory.AuthorID, fmt.Sprintf(expectReadEqualMessage, "story"))
			assert.Equal(t, story.GroupID, dbStory.GroupID, fmt.Sprintf(expectReadEqualMessage, "story"))
			assert.Equal(t, story.Content, dbStory.Content, fmt.Sprintf(expectReadEqualMessage, "story"))
		}
	})
}

func TestStoryDB(t *testing.T) {
	t.Run("cannot create without author", func(t *testing.T) {
		db, cleanUp := testutils.SetupDBConnection(t, dbConfig, migrationPath)
		defer cleanUp(t)

		group := Group{
			Name: "testGroup",
		}
		_ = CreateGroup(db, &group)

		err := db.Exec(`INSERT INTO "stories" 
		("created_at","updated_at","deleted_at","author_id","group_id","title","content","pin_order") 
		VALUES ('2023-08-08 22:17:28.085','2023-08-08 22:17:28.085',NULL,NULL,NULL,'','# Hi, This is a test story.',NULL)`).
			Error
		var pgerr *pgconn.PgError
		if assert.ErrorAs(t, err, &pgerr, "Expected error when creating story without Author ID") {
			assert.Equal(t, errNonNullable.Code, pgerr.Code)
		}
	})
}
