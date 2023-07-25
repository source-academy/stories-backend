package storyviews

import "github.com/source-academy/stories-backend/model"

type ListView struct {
	AuthorID uint   `json:"authorId"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

func ListFrom(stories []model.Story) []ListView {
	storiesListView := make([]ListView, len(stories))
	for i, story := range stories {
		storiesListView[i] = ListView{
			AuthorID: story.AuthorID,
			Title:    story.Title,
			Content:  story.Content,
		}
	}
	return storiesListView
}
