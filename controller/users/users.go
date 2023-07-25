package users

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/internal/database"
	apierrors "github.com/source-academy/stories-backend/internal/errors"
	"github.com/source-academy/stories-backend/model"
	userviews "github.com/source-academy/stories-backend/view/users"
)

func HandleList(w http.ResponseWriter, r *http.Request) error {
	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	users, err := model.GetAllUsers(db)
	if err != nil {
		logrus.Error(err)
		return err
	}

	controller.EncodeJSONResponse(w, userviews.ListFrom(users))
	return nil
}

func HandleRead(w http.ResponseWriter, r *http.Request) error {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return apierrors.ClientBadRequestError{
			Message: fmt.Sprintf("Invalid userID: %v", err),
		}
	}

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	user, err := model.GetUserByID(db, userID)
	if err != nil {
		logrus.Error(err)
		return err
	}
	controller.EncodeJSONResponse(w, userviews.SingleFrom(user))
	return nil
}
