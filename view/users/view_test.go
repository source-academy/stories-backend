package userviews

import (
	"testing"

	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
	"github.com/source-academy/stories-backend/model"
	"github.com/stretchr/testify/assert"
)

func TestSummaryFrom(t *testing.T) {
	t.Run("should work with empty user_group", func(t *testing.T) {
		user := model.User{
			Username: "",
		}

		summaryView := SummaryFrom(user, model.UserGroup{})
		assert.Equal(t, uint(0), summaryView.GroupID, "group id should be 0")
		assert.Equal(t, "", summaryView.GroupName, "group name should be empty")
		assert.Equal(t, groupenums.Role(""), summaryView.Role, "role should be empty")
	})
}
