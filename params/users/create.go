package userparams

import "github.com/source-academy/stories-backend/model"

type Create struct {
	GithubUsername string `json:"githubUsername"`
	GithubID       int    `json:"githubId"`
}

// TODO: Add some validation
func (params *Create) Validate() error {
	return nil
}

func (params *Create) ToModel() *model.User {
	return &model.User{
		GithubUsername: params.GithubUsername,
		GithubID:       params.GithubID,
	}
}
