package userviews

import "github.com/source-academy/stories-backend/model"

type View struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	LoginProvider string `json:"provider"`
}

func SingleFrom(user model.User) View {
	name := user.FullName
	if name == "" {
		name = user.Username
	}
	userView := View{
		ID:            user.ID,
		Name:          name,
		Username:      user.Username,
		LoginProvider: user.LoginProvider.ToString(),
	}
	return userView
}

type SummaryView struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func SummaryFrom(user model.User) SummaryView {
	name := user.FullName
	if name == "" {
		name = user.Username
	}
	userView := SummaryView{
		ID:   user.ID,
		Name: name,
	}
	return userView
}
