// controller/story.go
package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/source-academy/stories-backend/model"
)

func GetStories(w http.ResponseWriter, r *http.Request) {
	stories := model.GetAllStories()
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(stories); err != nil {
		log.Fatalln(err)
		panic(err)
	}
}

func CreateStory(w http.ResponseWriter, r *http.Request) {
	var story model.Story
	if err := json.NewDecoder(r.Body).Decode(&story); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model.CreateStory(story)
	w.WriteHeader(http.StatusCreated)
}
