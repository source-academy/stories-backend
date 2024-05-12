package userviews

import (
	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
	"github.com/source-academy/stories-backend/model"
)

type ListView struct {
	ID            uint            `json:"id"`
	Name          string          `json:"name"`
	Username      string          `json:"username"`
	LoginProvider string          `json:"provider"`
	Role          groupenums.Role `json:"role"`
}

func ListFrom(users []model.User, roles []groupenums.Role) []ListView {
	usersListView := make([]ListView, len(users))
	for i, user := range users {
		usersListView[i] = ListView{
			// Unlike other views, we do not fallback an empty name to
			// the username for the users' list view.
			ID:            user.ID,
			Name:          user.FullName,
			Username:      user.Username,
			LoginProvider: user.LoginProvider.ToString(),
			Role:          roles[i],
		}
	}
	return usersListView
}
