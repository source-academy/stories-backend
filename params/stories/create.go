package storyparams

import "github.com/source-academy/stories-backend/model"

type Create struct {
	AuthorID uint   `json:"authorId"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

// TODO: Add some validation
func (params *Create) Validate() error {
	return nil
}

func (params *Create) ToModel() *model.Story {
	return &model.Story{
		AuthorID: params.AuthorID,
		Title:    params.Title,
		Content:  params.Content,
	}
}
