package userpermissiongroups

import (
	"net/http"

	"github.com/source-academy/stories-backend/internal/auth"
)

type IsSelf struct {
	UserID uint
}

func (p IsSelf) IsAuthorized(r *http.Request) bool {
	userID, err := auth.GetUserIDFrom(r)
	if err != nil {
		return false
	}

	return p.UserID == uint(*userID)
}
