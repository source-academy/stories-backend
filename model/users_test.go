package model

import (
	"testing"
	// "time"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	t.Run("should return correct number of users", func(t *testing.T) {
		users := GetAllUsers()

		assert.Len(t, users, 2, "Expected number of users to be 2")
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("should create a new user", func(t *testing.T) {
		user := User{
			UserID:         uuid.New(),
			GithubUsername: "User 3",
			GithubID:       3,
			// CreatedAt:      time.Now(),
			// DeletedAt:      time.Now(),
			// UpdatedAt:      time.Now(),
		}

		CreateUser(user)

		users := GetAllUsers()

		assert.Len(t, users, 3, "Expected number of users to be 3")

		lastStory := users[len(users)-1]
		assert.Equal(t, 3, lastStory.UserID, "Expected user ID to be 3")
		assert.Equal(t, "User 3", lastStory.GithubUsername, "Expected Github Username to be 'User 3'")
		assert.Equal(t, 3, lastStory.GithubID, "Expected Github ID to be 3'")
		// You can add assertions for CreatedAt, DeletedAt, and UpdatedAt if necessary
	})
}
