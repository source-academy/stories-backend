package model

import "testing"

// can run test using : go test -v ./model

func TestGetAllStories(t *testing.T) {
	stories := GetAllStories()

	if len(stories) != 2 {
		t.Errorf("Expected number of stories to be 2, but got %d", len(stories))
	}
}

func TestCreateStory(t *testing.T) {
	story := Story{
		ID:     3,
		Title:  "Story 3",
		Author: "Username3",
	}

	CreateStory(story)

	stories := GetAllStories()

	if len(stories) != 3 {
		t.Errorf("Expected number of stories to be 3, but got %d", len(stories))
	}

	lastStory := stories[len(stories)-1]
	if lastStory.ID != 3 {
		t.Errorf("Expected story ID to be 3, but got %d", lastStory.ID)
	}

	if lastStory.Title != "Story 3" {
		t.Errorf("Expected story title to be 'Story 3', but got '%s'", lastStory.Title)
	}

	if lastStory.Author != "Username3" {
		t.Errorf("Expected story author to be 'Username3', but got '%s'", lastStory.Author)
	}
}
