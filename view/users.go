package view

type User struct {
	ID             uint   `json:"userId"`
	GithubUsername string `json:"githubUsername"`
	GithubID       int    `json:"githubId"`
}
