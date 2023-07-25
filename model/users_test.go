package model

import (
	"fmt"
	"testing"

	"github.com/source-academy/stories-backend/internal/config"
	"github.com/source-academy/stories-backend/internal/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const (
	expectCreateEqualMessage = "Expected last %s to be the one created"
	expectReadEqualMessage   = "Expected read %s to be the one previously created"
)

var conf, _ = config.LoadFromEnvironment()
var dbConfig *config.DatabaseConfig = conf.Database

func setupDBConnection(t *testing.T, dbConfig *config.DatabaseConfig) (*gorm.DB, func(*testing.T)) {
	// TODO: Create test DB

	// Connect to DB
	db, err := database.Connect(dbConfig)
	if err != nil {
		t.Error(err)
	}

	return db, func(t *testing.T) {
		database.Close(db)

		// TODO: Drop test DB
	}
}

// FIXME: Coupling with the other operations in the stories database
// func TestGetAllUsers(t *testing.T) {
// 	t.Run("should return correct initial number of users", func(t *testing.T) {
// 		db, cleanUp := setupDBConnection(t, dbConfig)
// 		defer cleanUp(t)

// 		db.Exec("DELETE FROM users")
// 		users := GetAllUsers(db)
// 		assert.Len(t, users, 0, "Expected initial number of users to be 0")

// 		user := User{
// 			Username: "testUsername",
// 			LoginProvider:       123,
// 		}
// 		CreateUser(db, &user)
// 		users = GetAllUsers(db)
// 		assert.Len(t, users, 1, "Expected number of users to be 1 after adding 1 user")
// 	})
// }

func TestCreateUser(t *testing.T) {
	t.Run("should increase the total user count", func(t *testing.T) {
		db, cleanUp := setupDBConnection(t, dbConfig)
		defer cleanUp(t)

		initialUsers, err := GetAllUsers(db)
		assert.Nil(t, err, "Expected no error when getting all users")

		user := User{
			Username:      "testUsername0",
			LoginProvider: 123,
		}
		err = CreateUser(db, &user)
		assert.Nil(t, err, "Expected no error when creating user")

		newUsers, err := GetAllUsers(db)
		assert.Nil(t, err, "Expected no error when getting all users")
		assert.Len(t, newUsers, len(initialUsers)+1, "Expected number of users to increase by 1")

		var lastUser User
		db.Model(&User{}).Last(&lastUser)

		assert.Equal(t, user.ID, lastUser.ID, "Expected the user ID to be updated")
		assert.Equal(t, user.Username, lastUser.Username, fmt.Sprintf(expectCreateEqualMessage, "user"))
		assert.Equal(t, user.LoginProvider, lastUser.LoginProvider, fmt.Sprintf(expectCreateEqualMessage, "user"))
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("should get the correct user", func(t *testing.T) {
		db, cleanUp := setupDBConnection(t, dbConfig)
		defer cleanUp(t)

		users := []*User{
			{Username: "testUsername1", LoginProvider: 123},
			{Username: "testUsername2", LoginProvider: 456},
			{Username: "testUsername3", LoginProvider: 789},
		}

		for _, userToAdd := range users {
			_ = CreateUser(db, userToAdd)
		}

		for _, user := range users {
			// FIXME: Don't use typecast
			dbUser, err := GetUserByID(db, int(user.ID))
			assert.Nil(t, err, "Expected no error when getting user by valid ID")
			assert.Equal(t, user.ID, dbUser.ID, fmt.Sprintf(expectReadEqualMessage, "user"))
			assert.Equal(t, user.Username, dbUser.Username, fmt.Sprintf(expectReadEqualMessage, "user"))
			assert.Equal(t, user.LoginProvider, dbUser.LoginProvider, fmt.Sprintf(expectReadEqualMessage, "user"))
		}
	})
}
