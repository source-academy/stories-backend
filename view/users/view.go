package userviews

import (
	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
	"github.com/source-academy/stories-backend/model"
)

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
	// The following three field can be nested when we have many groups in a course
	GroupID   uint            `json:"group"`
	GroupName string          `json:"groupName"`
	Role      groupenums.Role `json:"role"`
}

func SummaryFrom(user model.User, userGroup model.UserGroup) SummaryView {
	name := user.FullName
	if name == "" {
		name = user.Username
	}
	userView := SummaryView{
		ID:        user.ID,
		Name:      name,
		GroupID:   userGroup.GroupID,
		GroupName: userGroup.Group.Name,
		Role:      userGroup.Role,
	}
	return userView
}
