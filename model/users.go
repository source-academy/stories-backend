package model

import "time"

type User struct {
	UserID         int       `json:"user_id"`
	GithubUsername string    `json:"github_username"`
	GithubID       int       `json:"github_id"`
	CreatedAt      time.Time `json:"created_at"`
	DeletedAt      time.Time `json:"deleted_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

var users = []User{
	{UserID: 1, GithubUsername: "User 1", GithubID: 1, CreatedAt: time.Now(), DeletedAt: time.Now(), UpdatedAt: time.Now()},
	{UserID: 2, GithubUsername: "User 2", GithubID: 2, CreatedAt: time.Now(), DeletedAt: time.Now(), UpdatedAt: time.Now()},
}

func GetAllUsers() []User {
	return users
}

func GetUserByID(userID int) *User {
	for i, user := range users {
		if user.UserID == userID {
			return &users[i]
		}
	}
	return nil
}

func CreateUser(user User) {
	users = append(users, user)
}
