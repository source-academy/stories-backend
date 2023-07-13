package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/model"
	"github.com/source-academy/stories-backend/view"
)

func EncodeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(data); err != nil {
		logrus.Errorln(err)
		panic(err)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := model.GetAllUsers()
	EncodeJSONResponse(w, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}
	user := model.GetUserByID(userID)
	EncodeJSONResponse(w, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user view.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	model.CreateUser(user)
	EncodeJSONResponse(w, &user)
	w.WriteHeader(http.StatusCreated)
}
