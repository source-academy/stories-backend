package storyparams

import (
	"github.com/source-academy/stories-backend/model"
)

type Update struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	PinOrder *int   `json:"pinOrder"`
}

func (params *Update) Validate() error {
	// Extra params won't do anything, e.g. authorID can't be changed.
	// TODO: Error on extra params?
	return nil
}

func (params *Update) ToModel() *model.Story {
	return &model.Story{
		Title:    params.Title,
		Content:  params.Content,
		PinOrder: params.PinOrder,
	}
}
