package users

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/model"
	"github.com/source-academy/stories-backend/view"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := model.GetAllUsers()
	controller.EncodeJSONResponse(w, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}
	user := model.GetUserByID(userID)
	controller.EncodeJSONResponse(w, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user view.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	model.CreateUser(user)
	controller.EncodeJSONResponse(w, &user)
	w.WriteHeader(http.StatusCreated)
}
