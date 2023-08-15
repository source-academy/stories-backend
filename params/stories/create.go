package storyparams

import (
	"github.com/source-academy/stories-backend/model"
)

type Create struct {
	AuthorID uint   `json:"authorId"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	PinOrder *int   `json:"pinOrder"`
}

// TODO: Add some validation
func (params *Create) Validate() error {
	return nil
}

func (params *Create) ToModel(associatedGroupID *uint) *model.Story {
	return &model.Story{
		AuthorID: params.AuthorID,
		GroupID:  associatedGroupID,
		Title:    params.Title,
		Content:  params.Content,
		PinOrder: params.PinOrder,
	}
}
