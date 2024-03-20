package model

import (
	"fmt"

	"github.com/source-academy/stories-backend/internal/database"
	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StoryStatus int

const (
    Draft StoryStatus = iota
    Published
)

type Story struct {
    gorm.Model
    AuthorID   uint
    Author     User
    GroupID    *uint
    Group      Group
    Title      string
    Content    string
    PinOrder   *int
    Status     StoryStatus
}

// Passing nil to omit the filtering and get all stories
// TODO: Use nullable types instead
func GetAllStoriesInGroup(db *gorm.DB, groupID *uint) ([]Story, error) {
	var stories []Story
	err := db.
		// FIXME: Handle nil case properly
		Where(Story{GroupID: groupID}).
		Preload(clause.Associations).
		// TODO: Abstract out the sorting logic
		Order("pin_order ASC NULLS LAST, title ASC, content ASC").
		Find(&stories).
		Error
	if err != nil {
		return stories, database.HandleDBError(err, "story")
	}
	return stories, nil
}

func GetAllPublishedStories(db *gorm.DB, groupID *uint) ([]Story, error) {
	var stories []Story
	err := db.
		// FIXME: Handle nil case properly
		Where(Story{GroupID: groupID, Status: Published}).
		Preload(clause.Associations).
		// TODO: Abstract out the sorting logic
		Order("pin_order ASC NULLS LAST, title ASC, content ASC").
		Find(&stories).
		Error
	if err != nil {
		return stories, database.HandleDBError(err, "story")
	}
	return stories, nil
}

func GetAllDraftStories(db *gorm.DB, groupID *uint) ([]Story, error) {
	var stories []Story
	err := db.
		// FIXME: Handle nil case properly
		Where(Story{GroupID: groupID, Status: Draft}).
		Preload(clause.Associations).
		// TODO: Abstract out the sorting logic
		Order("pin_order ASC NULLS LAST, title ASC, content ASC").
		Find(&stories).
		Error
	if err != nil {
		return stories, database.HandleDBError(err, "story")
	}
	return stories, nil
}

func GetStoryByID(db *gorm.DB, id int) (Story, error) {
	var story Story
	err := db.
		Preload(clause.Associations).
		First(&story, id).
		Error
	if err != nil {
		return story, database.HandleDBError(err, "story")
	}
	return story, nil
}

func CreateStory(db *gorm.DB, story *Story) error {
	// Check author's role
	role, err := GetUserRoleByID(db, story.AuthorID)
	if err != nil {
		return fmt.Errorf("failed to get user role: %w", err)
	}

	// Set story status based on author's role
	if role == groupenums.RoleStandard {
		story.Status = Draft
	} else {
		story.Status = Published
	}

	// Create the story in the database
	if err := db.
		Preload(clause.Associations).
		Create(story).
		// Get associated Author.
		First(story).
		Error; err != nil {
		return fmt.Errorf("failed to create story: %w", err)
	}

	return nil
}

func UpdateStory(db *gorm.DB, storyID int, newStory *Story) error {
	// TODO: Possible restore functionality for soft-deleted stories?
	err := db.
		Transaction(func(tx *gorm.DB) error {
			var originalStory Story
			err := tx.
				Preload(clause.Associations).
				Where("id = ?", storyID).
				First(&originalStory).
				Error
			if err != nil {
				return database.HandleDBError(err, "story")
			}

			// Handle nullable fields when null
			if newStory.PinOrder == nil {
				err = tx.
					Preload(clause.Associations).
					Where("id = ?", storyID).
					Model(newStory).
					Update("pin_order", gorm.Expr("NULL")).
					Error
				if err != nil {
					return database.HandleDBError(err, "story")
				}
			}
			// Update remaining fields
			err = tx.
				Preload(clause.Associations).
				Where("id = ?", storyID).
				Updates(newStory).
				First(newStory).
				Error
			if err != nil {
				return database.HandleDBError(err, "story")
			}

			return nil
		})
	if err != nil {
		return database.HandleDBError(err, "story")
	}

	return nil
}

func DeleteStory(db *gorm.DB, storyID int) (Story, error) {
	var story Story
	err := db.
		Preload(clause.Associations).
		Where("id = ?", storyID).
		First(&story). // store the value to be returned
		Delete(&story).
		Error
	if err != nil {
		return story, database.HandleDBError(err, "story")
	}
	return story, nil
}
