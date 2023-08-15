package storypermissiongroups

import (
	"net/http"

	"github.com/source-academy/stories-backend/internal/auth"
	"github.com/source-academy/stories-backend/internal/database"
	"github.com/source-academy/stories-backend/model"
)

type IsAuthorOf struct {
	StoryID uint
}

func (p IsAuthorOf) IsAuthorized(r *http.Request) bool {
	userID, err := auth.GetUserIDFrom(r)
	if err != nil {
		return false
	}

	db, err := database.GetDBFrom(r)
	if err != nil {
		return false
	}

	story, err := model.GetStoryByID(db, int(p.StoryID))
	if err != nil {
		return false
	}

	return story.AuthorID == uint(*userID)
}
