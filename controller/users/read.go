package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/internal/auth"
	"github.com/source-academy/stories-backend/internal/database"
	apierrors "github.com/source-academy/stories-backend/internal/errors"
	"github.com/source-academy/stories-backend/model"
	userviews "github.com/source-academy/stories-backend/view/users"
)

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

func HandleReadSelf(w http.ResponseWriter, r *http.Request) error {
	userID, err := auth.GetUserIDFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	user, err := model.GetUserByID(db, *userID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	var userGroup model.UserGroup
	courseIDStr := r.URL.Query().Get("course")
	if courseIDStr != "" {
		courseID, err := strconv.Atoi(courseIDStr)
		if err != nil {
			logrus.Error(err)
			return err
		}

		courseGroup, err := model.GetGroupByCourseID(db, courseID)
		if err != nil {
			logrus.Error(err)
			return err
		}

		userGroup, err = model.GetUserGroupByID(db, user.ID, courseGroup.ID)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	controller.EncodeJSONResponse(w, userviews.SummaryFrom(user, userGroup))
	return nil
}
