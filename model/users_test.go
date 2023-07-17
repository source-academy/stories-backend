package model

import (
	"testing"

	"github.com/source-academy/stories-backend/internal/config"
	"github.com/source-academy/stories-backend/internal/database"
	dbutils "github.com/source-academy/stories-backend/internal/utils/db"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const (
	expectCreateEqualMessage = "Expected last user to be the one created"
	expectReadEqualMessage   = "Expected read user to be the one previously created"
)

var dbConfig *config.DatabaseConfig = &config.DatabaseConfig{
	TimeZone:     dbutils.DB_DEFAULT_TIMEZONE,
	Host:         dbutils.DB_DEFAULT_HOSTNAME,
	User:         dbutils.DB_DEFAULT_USER,
	Password:     dbutils.DB_DEFAULT_PASSWORD,
	Port:         dbutils.DB_DEFAULT_PORT,
	DatabaseName: dbutils.DB_DEFAULT_NAME,
}

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

func TestGetAllUsers(t *testing.T) {
	t.Run("should return correct initial number of users", func(t *testing.T) {
		db, cleanUp := setupDBConnection(t, dbConfig)
		defer cleanUp(t)

		users := GetAllUsers(db)
		assert.Len(t, users, 2, "Expected initial number of users to be 0")
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("should increase the total user count", func(t *testing.T) {
		db, cleanUp := setupDBConnection(t, dbConfig)
		defer cleanUp(t)

		initialUsers := GetAllUsers(db)

		user := User{
			GithubUsername: "testUsername",
			GithubID:       123,
		}
		CreateUser(db, &user)

		newUsers := GetAllUsers(db)
		assert.Len(t, newUsers, len(initialUsers)+1, "Expected number of users to increase by 1")

		var lastUser User
		db.Model(&User{}).Last(&lastUser)

		assert.Equal(t, user.ID, lastUser.ID, expectCreateEqualMessage)
		assert.Equal(t, user.GithubUsername, lastUser.GithubUsername, expectCreateEqualMessage)
		assert.Equal(t, user.GithubID, lastUser.GithubID, expectCreateEqualMessage)
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("should get the correct user", func(t *testing.T) {
		db, cleanUp := setupDBConnection(t, dbConfig)
		defer cleanUp(t)

		users := []User{
			{GithubUsername: "testUsername1", GithubID: 123},
			{GithubUsername: "testUsername2", GithubID: 456},
			{GithubUsername: "testUsername3", GithubID: 789},
		}

		for _, userToAdd := range users {
			CreateUser(db, &userToAdd)
		}

		for _, user := range users {
			// FIXME: Don't use typecast
			dbUser := GetUserByID(db, int(user.ID))
			assert.Equal(t, user.ID, dbUser.ID, expectReadEqualMessage)
			assert.Equal(t, user.GithubUsername, dbUser.GithubUsername, expectReadEqualMessage)
			assert.Equal(t, user.GithubID, dbUser.GithubID, expectReadEqualMessage)
		}
	})
}
