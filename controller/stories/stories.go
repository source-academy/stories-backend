package stories

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/model"
	storyviews "github.com/source-academy/stories-backend/view/stories"
)

func GetStories(w http.ResponseWriter, r *http.Request) {
	stories := model.GetAllStories()
	controller.EncodeJSONResponse(w, stories)
}

func GetStory(w http.ResponseWriter, r *http.Request) {
	storyIDStr := chi.URLParam(r, "storyID")
	storyID, err := strconv.Atoi(storyIDStr)
	if err != nil {
		http.Error(w, "Invalid storyID", http.StatusBadRequest)
		return
	}
	story := model.GetStoryByID(storyID)
	controller.EncodeJSONResponse(w, story)
}

func CreateStory(w http.ResponseWriter, r *http.Request) {
	var story storyviews.View
	if err := json.NewDecoder(r.Body).Decode(&story); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model.CreateStory(story)
	controller.EncodeJSONResponse(w, &story)
	w.WriteHeader(http.StatusCreated)
}
