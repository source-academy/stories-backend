package userviews

import "github.com/source-academy/stories-backend/model"

type ListView struct {
	Username      string `json:"username"`
	LoginProvider string `json:"provider"`
}

func ListFrom(users []model.User) []ListView {
	usersListView := make([]ListView, len(users))
	for i, user := range users {
		usersListView[i] = ListView{
			Username:      user.Username,
			LoginProvider: user.LoginProvider.String(),
		}
	}
	return usersListView
}
