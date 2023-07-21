package userviews

import "github.com/source-academy/stories-backend/model"

type View struct {
	ID            uint   `json:"id"`
	Username      string `json:"username"`
	LoginProvider string `json:"provider"`
}

func SingleFrom(user model.User) View {
	userView := View{
		ID:            user.ID,
		Username:      user.Username,
		LoginProvider: user.LoginProvider.ToString(),
	}
	return userView
}
