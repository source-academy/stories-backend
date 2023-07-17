package storyviews

import "github.com/source-academy/stories-backend/model"

type View struct {
	ID           uint   `json:"storyId"`
	AuthorID     uint   `json:"authorId"`
	StoryContent string `json:"storyContent"`
}

func SingleFrom(story model.Story) View {
	storyView := View{
		ID:           story.ID,
		AuthorID:     story.AuthorID,
		StoryContent: story.StoryContent,
	}
	return storyView
}
