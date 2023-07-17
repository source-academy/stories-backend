package userviews

import "github.com/source-academy/stories-backend/model"

type View struct {
	ID             uint   `json:"userId"`
	GithubUsername string `json:"githubUsername"`
	GithubID       int    `json:"githubId"`
}

func SingleFrom(user model.User) View {
	userView := View{
		ID:             user.ID,
		GithubUsername: user.GithubUsername,
		GithubID:       user.GithubID,
	}
	return userView
}
