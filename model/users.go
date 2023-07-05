package model

type User struct {
	UserID         int    `json:"user_id"`
	GithubUsername string `json:"github_username"`
	GithubID       int    `json:"github_ID"`
}

var users = []User{
	{UserID: 1, GithubUsername: "User 1", GithubID: 1},
	{UserID: 2, GithubUsername: "User 2", GithubID: 2},
}

func GetAllUsers() []User {
	return users
}

func CreateUser(user User) {
	users = append(users, user)
}
