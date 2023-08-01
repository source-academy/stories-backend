package model

import (
	"fmt"
	"math/rand"
	"testing"

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
			dbUserGroup, err := GetUserGroupByID(db, int(users[idx].ID), int(group.ID))
			assert.Nil(t, err, "Expected no error when getting userGroup with valid IDs")
			assert.Equal(t, userGroup.Role, dbUserGroup.Role, fmt.Sprintf(expectCreateEqualMessage, "usergroup"))
			assert.Equal(t, userGroup.UserID, dbUserGroup.UserID, fmt.Sprintf(expectCreateEqualMessage, "usergroup"))
			assert.Equal(t, userGroup.User.Username, dbUserGroup.User.Username, fmt.Sprintf(expectCreateEqualMessage, "usergroup"))
			assert.Equal(t, userGroup.Group.ID, dbUserGroup.Group.ID, fmt.Sprintf(expectCreateEqualMessage, "usergroup"))
		}

		_, err := GetUserGroupByID(db, int(users[1].ID+1), int(group.ID))
		// is this the correct bahaviour? calling to a model function returns a api error
		// assert.ErrorIs(t, err, gorm.ErrRecordNotFound, "Expected error when getting userGroup with invalid ID")
		assert.ErrorIs(t, err, apierrors.ClientNotFoundError{Message: "Cannot find requested userGroup."}, "Expected error when getting userGroup with invalid ID")
	})
}
