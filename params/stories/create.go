package storyparams

import "github.com/source-academy/stories-backend/model"

type Create struct {
	AuthorID     uint   `json:"authorId"`
	StoryContent string `json:"storyContent"`
}

// TODO: Add some validation
func (params *Create) Validate() error {
	return nil
}

func (params *Create) ToModel() *model.Story {
	return &model.Story{
		AuthorID:     params.AuthorID,
		StoryContent: params.StoryContent,
	}
}
