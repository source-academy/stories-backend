package view

type User struct {
	UserID         int    `json:"userId"`
	GithubUsername string `json:"githubUsername"`
	GithubID       int    `json:"githubId"`
}
