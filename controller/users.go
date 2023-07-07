package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/model"

	"gorm.io/gorm"
)

var DB *gorm.DB

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// users := model.GetAllUsers()
	var users []model.User
	if err := DB.Select("user_id, github_username, github_id").Find(&users).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(users); err != nil {
		logrus.Errorln(err)
		panic(err)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}
	user := model.GetUserByID(userID)
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(user); err != nil {
		logrus.Errorln(err)
		panic(err)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// model.CreateUser(user)
	if err := DB.Exec("INSERT INTO users (user_id, github_username, github_id) VALUES ($1, $2, $3)", user.UserID, user.GithubUsername, user.GithubID).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
