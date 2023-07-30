package storyviews

import (
	"github.com/source-academy/stories-backend/model"
	userviews "github.com/source-academy/stories-backend/view/users"
)

type View struct {
	ID         uint   `json:"id"`
	AuthorID   uint   `json:"authorId"`
	AuthorName string `json:"authorName"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Pinned     bool   `json:"isPinned"`
	PinOrder   *int   `json:"pinOrder"`
}

func SingleFrom(story model.Story) View {
	author := userviews.SingleFrom(story.Author)
	storyView := View{
		ID:         story.ID,
		AuthorID:   story.AuthorID,
		AuthorName: author.Name,
		Title:      story.Title,
		Content:    story.Content,
		Pinned:     story.PinOrder != nil,
		PinOrder:   story.PinOrder,
	}
	return storyView
}
