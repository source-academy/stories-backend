package model

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
	userenums "github.com/source-academy/stories-backend/internal/enums/users"
	apierrors "github.com/source-academy/stories-backend/internal/errors"
	"github.com/source-academy/stories-backend/internal/testutils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
)

func TestCreateUserGroup(t *testing.T) {
	t.Run("succeeds: should increase the total usergroup count", func(t *testing.T) {
		db, cleanUp := testutils.SetupDBConnection(t, dbConfig, migrationPath)
		defer cleanUp(t)

		// We need to first create a user and a group due to the foreign key constraint
		user := User{
			Username:      "testUser1",
			LoginProvider: userenums.LoginProvider(rand.Int31()),
		}
		_ = CreateUser(db, &user)

		group := Group{
			Name: "testGroup",
		}
		_ = CreateGroup(db, &group)

		userGroup := UserGroup{
			User:  user,
			Group: group,
			Role:  groupenums.RoleStandard,
		}
		err := CreateUserGroup(db, &userGroup)
		assert.Nil(t, err, "Expected no error when creating usergroup")

		var lastUserGroup UserGroup
		db.Preload(clause.Associations). // same as Preload("User").Preload("Group"). to preload all
							Last(&lastUserGroup)

		assert.Equal(t, userGroup.ID, lastUserGroup.ID, "Expected the usergroup ID to be updated")
		assert.Equal(t, userGroup.Role, lastUserGroup.Role, fmt.Sprintf(expectCreateEqualMessage, "usergroup"))
		assert.Equal(t, userGroup.UserID, lastUserGroup.UserID, fmt.Sprintf(expectCreateEqualMessage, "usergroup"))
		assert.Equal(t, userGroup.User.Username, lastUserGroup.User.Username, fmt.Sprintf(expectCreateEqualMessage, "usergroup"))
		assert.Equal(t, userGroup.Group.ID, lastUserGroup.Group.ID, fmt.Sprintf(expectCreateEqualMessage, "usergroup"))
	})

	t.Run("failure: should not create twice with user_group index", func(t *testing.T) {
		db, cleanUp := testutils.SetupDBConnection(t, dbConfig, migrationPath)
		defer cleanUp(t)

		// We need to first create a user and a group due to the foreign key constraint
		user := User{
			Username:      "testUser1",
			LoginProvider: userenums.LoginProvider(rand.Int31()),
		}
		_ = CreateUser(db, &user)

		group := Group{
			Name: "testGroup",
		}
		_ = CreateGroup(db, &group)

		userGroup := UserGroup{
			User:  user,
			Group: group,
			Role:  groupenums.RoleStandard,
		}
		err := CreateUserGroup(db, &userGroup)
		assert.Nil(t, err, "Expected no error when creating usergroup for the first time")

		userGroup2 := UserGroup{
			User:  user,
			Group: group,
			Role:  groupenums.RoleStandard,
		}
		err = CreateUserGroup(db, &userGroup2)
		assert.Error(t, err, "Expected error for violating user_group unique index")
	})

	t.Run("failure: should not create with missing user or group", func(t *testing.T) {
		db, cleanUp := testutils.SetupDBConnection(t, dbConfig, migrationPath)
		defer cleanUp(t)

		// We need to first create a user and a group due to the foreign key constraint
		user := User{
			Username:      "testUser1",
			LoginProvider: userenums.LoginProvider(rand.Int31()),
		}
		_ = CreateUser(db, &user)

		group := Group{
			Name: "testGroup",
		}
		_ = CreateGroup(db, &group)

		userGroup := UserGroup{
			User: user,
			Role: groupenums.RoleStandard,
		}
		err := CreateUserGroup(db, &userGroup)
		var pgerr *pgconn.PgError
		if assert.ErrorAs(t, err, &pgerr, "Expected error when creating story without Author ID") {
			assert.Equal(t, errInvalidForeignKey.Code, pgerr.Code)
		}
		// assert.Error(t, err, "Expected error when creating usergroup without group")

		userGroup2 := UserGroup{
			Group: group,
			Role:  groupenums.RoleStandard,
		}
		err = CreateUserGroup(db, &userGroup2)
		var pgerr2 *pgconn.PgError
		if assert.ErrorAs(t, err, &pgerr2, "Expected error when creating story without Author ID") {
			assert.Equal(t, errInvalidForeignKey.Code, pgerr2.Code)
		}
		// assert.Error(t, err, "Expected error when creating usergroup without user")
	})
}

func TestGetUserGroupByID(t *testing.T) {
	// t.Skip("skip")
	t.Run("should get the correct story", func(t *testing.T) {
		db, cleanUp := testutils.SetupDBConnection(t, dbConfig, migrationPath)
		defer cleanUp(t)

		// We need to first create a user and a group due to the foreign key constraint
		group := Group{
			Name: "testGroup",
		}
		_ = CreateGroup(db, &group)

		users := []*User{
			{Username: "testUser1",
				LoginProvider: userenums.LoginProvider(rand.Int31())},
			{Username: "testUser2",
				LoginProvider: userenums.LoginProvider(rand.Int31())},
		}
		usergroups := []*UserGroup{}

		for _, userToAdd := range users {
			_ = CreateUser(db, userToAdd)

			userGroup := UserGroup{
				User:  *userToAdd,
				Group: group,
				Role:  groupenums.RoleStandard,
			}
			usergroups = append(usergroups, &userGroup)
			err := CreateUserGroup(db, &userGroup)
			assert.Nil(t, err, "Expected no error when creating usergroup")
		}

		for idx, userGroup := range usergroups {
			// FIXME: Don't use typecast
			dbUserGroup, err := GetUserGroupByID(db, users[idx].ID, group.ID)
			assert.Nil(t, err, "Expected no error when getting userGroup with valid IDs")
			assert.Equal(t, userGroup.Role, dbUserGroup.Role, fmt.Sprintf(expectCreateEqualMessage, "usergroup"))
			assert.Equal(t, userGroup.UserID, dbUserGroup.UserID, fmt.Sprintf(expectCreateEqualMessage, "usergroup"))
			assert.Equal(t, userGroup.User.Username, dbUserGroup.User.Username, fmt.Sprintf(expectCreateEqualMessage, "usergroup"))
			assert.Equal(t, userGroup.Group.ID, dbUserGroup.Group.ID, fmt.Sprintf(expectCreateEqualMessage, "usergroup"))
		}

		_, err := GetUserGroupByID(db, users[1].ID+1, group.ID)
		// is this the correct bahaviour? calling to a model function returns a api error
		// assert.ErrorIs(t, err, gorm.ErrRecordNotFound, "Expected error when getting userGroup with invalid ID")
		assert.ErrorIs(t, err, apierrors.ClientNotFoundError{Message: "Cannot find requested userGroup."}, "Expected error when getting userGroup with invalid ID")
	})
}

func TestUserGroupDB(t *testing.T) {
	t.Run("cannot create without user", func(t *testing.T) {
		db, cleanUp := testutils.SetupDBConnection(t, dbConfig, migrationPath)
		defer cleanUp(t)

		// We need to first create a user and a group due to the foreign key constraint
		user := User{
			Username:      "testUser1",
			LoginProvider: userenums.LoginProvider(rand.Int31()),
		}
		_ = CreateUser(db, &user)

		group := Group{
			Name: "testGroup",
		}
		_ = CreateGroup(db, &group)

		err := db.Exec(`INSERT INTO "user_groups" 
			("created_at","updated_at","deleted_at","user_id","group_id","role") 
			VALUES ('2023-08-08 23:44:01.417','2023-08-08 23:44:01.417',NULL,NULL,1,'member')`).
			Error
		var pgerr *pgconn.PgError
		if assert.ErrorAs(t, err, &pgerr, "Expected error when creating story without Author ID") {
			assert.Equal(t, errNonNullable.Code, pgerr.Code)
		}
	})

	t.Run("cannot create without group", func(t *testing.T) {
		db, cleanUp := testutils.SetupDBConnection(t, dbConfig, migrationPath)
		defer cleanUp(t)

		// We need to first create a user and a group due to the foreign key constraint
		user := User{
			Username:      "testUser1",
			LoginProvider: userenums.LoginProvider(rand.Int31()),
		}
		_ = CreateUser(db, &user)

		group := Group{
			Name: "testGroup",
		}
		_ = CreateGroup(db, &group)

		err := db.Exec(`INSERT INTO "user_groups" 
			("created_at","updated_at","deleted_at","user_id","group_id","role") 
			VALUES ('2023-08-08 23:44:01.417','2023-08-08 23:44:01.417',NULL,1,NULL,'member')`).
			Error
		var pgerr *pgconn.PgError
		if assert.ErrorAs(t, err, &pgerr, "Expected error when creating story without Author ID") {
			assert.Equal(t, errNonNullable.Code, pgerr.Code)
		}
	})
}
