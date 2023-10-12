package userviews

import "github.com/source-academy/stories-backend/model"

type ListView struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	LoginProvider string `json:"provider"`
}

func ListFrom(users []model.User) []ListView {
	usersListView := make([]ListView, len(users))
	for i, user := range users {
		usersListView[i] = ListView{
			// Unlike other views, we do not fallback an empty name to
			// the username for the users' list view.
			ID:            user.ID,
			Name:          user.FullName,
			Username:      user.Username,
			LoginProvider: user.LoginProvider.ToString(),
		}
	}
	return usersListView
}
