package storyviews

type View struct {
	ID           uint   `json:"storyId"`
	AuthorID     uint   `json:"authorId"`
	StoryContent string `json:"storyContent"`
}
