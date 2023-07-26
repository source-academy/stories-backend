package storyviews

import (
	"github.com/source-academy/stories-backend/model"
	userviews "github.com/source-academy/stories-backend/view/users"
)

type ListView struct {
	AuthorID   uint   `json:"authorId"`
	AuthorName string `json:"authorName"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Pinned     bool   `json:"isPinned"`
}

func ListFrom(stories []model.Story) []ListView {
	storiesListView := make([]ListView, len(stories))
	for i, story := range stories {
		author := userviews.SingleFrom(story.Author)
		storiesListView[i] = ListView{
			AuthorID:   story.AuthorID,
			AuthorName: author.Name,
			Title:      story.Title,
			Content:    story.Content,
			Pinned:     story.PinOrder != nil,
		}
	}
	return storiesListView
}
