package storyviews

import "github.com/source-academy/stories-backend/model"

type ListView struct {
	AuthorID     uint   `json:"authorId"`
	StoryContent string `json:"storyContent"`
}

func ListFrom(stories []model.Story) []ListView {
	storiesListView := make([]ListView, len(stories))
	for i, story := range stories {
		storiesListView[i] = ListView{
			AuthorID:     story.AuthorID,
			StoryContent: story.StoryContent,
		}
	}
	return storiesListView
}
