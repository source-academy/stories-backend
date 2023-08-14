package model

import (
	"testing"

	"github.com/source-academy/stories-backend/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestGetGroupByCourseID(t *testing.T) {
	t.Run("should get the correct group", func(t *testing.T) {
		db, cleanUp := testutils.SetupDBConnection(t, dbConfig, migrationPath)
		defer cleanUp(t)

		group := Group{
			Name: "testGroup",
		}
		_ = CreateGroup(db, &group)

		err := db.Exec(`INSERT INTO "course_groups" 
						("course_id","group_id") 
						VALUES (42, ?)`, group.ID).Error
		assert.Nil(t, err, "Expected no error when inserting course group")

		dbGroup, err := GetGroupByCourseID(db, 42)
		assert.Nil(t, err, "Expected no error when getting course group")
		assert.Equal(t, group.ID, dbGroup.ID, "expect to get the correct group.")
		assert.Equal(t, group.Name, dbGroup.Name, "expect to get the correct group.")
	})
}
