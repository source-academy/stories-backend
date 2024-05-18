package model

import (
	"github.com/source-academy/stories-backend/internal/database"
	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StoryStatus int

const (
	Draft StoryStatus = iota
	Pending
	Rejected
	Published
)

type Story struct {
	gorm.Model
	AuthorID      uint
	Author        User
	GroupID       *uint // null means this is a public story
	Group         Group
	Title         string
	Content       string
	PinOrder      *int // nil if not pinned
	Status        StoryStatus
	StatusMessage *string
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
		Where("status = ?", int(Published)).
		Where("group_id = ?", groupID).
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

func GetAllPendingStories(db *gorm.DB, groupID *uint) ([]Story, error) {
	var stories []Story
	err := db.
		Where("status = ?", int(Pending)).
		Where("group_id = ?", groupID).
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

func GetAllStoriesByStatus(db *gorm.DB, groupID *uint, status StoryStatus) ([]Story, error) {
	var stories []Story
	err := db.
		Where("status = ?", int(status)).
		Where("group_id = ?", groupID).
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

func GetAllAuthorStoriesByStatus(db *gorm.DB, groupID *uint, userID *int, status StoryStatus) ([]Story, error) {
	var stories []Story
	err := db.
		Where("status = ?", int(status)).
		Where("group_id = ?", groupID).
		Where("author_id = ?", userID).
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

func (s *Story) create(tx *gorm.DB) *gorm.DB {
	return tx.
		Preload(clause.Associations).
		Create(s).
		// Get associated Author. See
		// https://github.com/go-gorm/gen/issues/618 on why
		// a separate .First() is needed.
		First(s)
}

func CreateStory(db *gorm.DB, story *Story) error {
	// Check author's role
	role, _ := GetUserRoleByID(db, story.AuthorID)
	// Based on the TestCreateStory, "can create without group" seems to be the desired behaviour
	// No group means no userGroup, which means no role, so an error shouldn't be thrown
	// Set story status based on author's role
	if !groupenums.IsRoleGreaterThan(role, groupenums.RoleStandard) {
		story.Status = Draft
	} else {
		story.Status = Published
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(story).Error; err != nil {
			return err // Return the error directly
		}
		return nil
	})
	if err != nil {
		return database.HandleDBError(err, "story")
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

func (s *Story) delete(tx *gorm.DB, storyID uint) *gorm.DB {
	return tx.
		Preload(clause.Associations).
		Where("id = ?", storyID).
		First(s). // store the value to be returned
		Delete(s)
}

func DeleteStory(db *gorm.DB, storyID int) (Story, error) {
	var story Story
	err := db.Transaction(func(tx *gorm.DB) error {
		return story.delete(tx, uint(storyID)).Error
	})
	if err != nil {
		return story, database.HandleDBError(err, "story")
	}
	return story, nil
}
