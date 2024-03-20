package storyparams

import (
    "github.com/source-academy/stories-backend/model"
)

type Publish struct {
    IsPublished bool `json:"boolean"`
}

// Validate validates the Publish params.
func (params *Publish) Validate() error {
    // Validation logic can be added here if needed.
    return nil
}

// ToModel converts Publish params to a Story model.
func (params *Publish) ToModel() *model.Story {
	if (params.IsPublished) {
		return &model.Story{
			Status: model.Published,
		}
	} else {
		return &model.Story{
			Status: model.Draft,
		}
	}
}