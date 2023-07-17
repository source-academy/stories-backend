package userviews

import "github.com/source-academy/stories-backend/model"

type ListView struct {
	GithubUsername string `json:"githubUsername"`
	GithubID       int    `json:"githubId"`
}

func ListFrom(users []model.User) []ListView {
	usersListView := make([]ListView, len(users))
	for i, user := range users {
		usersListView[i] = ListView{
			GithubUsername: user.GithubUsername,
			GithubID:       user.GithubID,
		}
	}
	return usersListView
}
