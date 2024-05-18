package storyparams

import (
	"fmt"
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
	if params.AuthorID == 0 {
		return fmt.Errorf("authorId is required and must be non-zero")
	}
	if params.Title == "" {
		return fmt.Errorf("title is required and cannot be empty")
	}
	if params.Content == "" {
		return fmt.Errorf("content is required and cannot be empty")
	}
	if params.PinOrder != nil && *params.PinOrder < 0 {
		return fmt.Errorf("pinOrder, if set, must be non-negative")
	}
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
