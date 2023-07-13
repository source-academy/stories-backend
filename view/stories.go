package view

type Story struct {
	StoryID      int    `json:"storyId"`
	AuthorID     int    `json:"authorId"`
	StoryContent string `json:"storyContent"`
}
