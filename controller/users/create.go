package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/internal/auth"
	"github.com/source-academy/stories-backend/internal/database"
	apierrors "github.com/source-academy/stories-backend/internal/errors"
	userpermissiongroups "github.com/source-academy/stories-backend/internal/permissiongroups/users"
	"github.com/source-academy/stories-backend/model"
	userparams "github.com/source-academy/stories-backend/params/users"
	userviews "github.com/source-academy/stories-backend/view/users"
)

func HandleCreate(w http.ResponseWriter, r *http.Request) error {
	err := auth.CheckPermissions(r, userpermissiongroups.Create())
	if err != nil {
		logrus.Error(err)
		return apierrors.ClientForbiddenError{
			Message: fmt.Sprintf("Error creating user: %v", err),
		}
	}

	var params userparams.Create
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		e, ok := err.(*json.UnmarshalTypeError)
		if !ok {
			logrus.Error(err)
			return apierrors.ClientBadRequestError{
				Message: fmt.Sprintf("Bad JSON parsing: %v", err),
			}
		}

		// TODO: Investigate if we should use errors.Wrap instead
		return apierrors.ClientUnprocessableEntityError{
			Message: fmt.Sprintf("Invalid JSON format: %s should be a %s.", e.Field, e.Type),
		}
	}

	err = params.Validate()
	if err != nil {
		logrus.Error(err)
		return apierrors.ClientUnprocessableEntityError{
			Message: fmt.Sprintf("JSON validation failed: %v", err),
		}
	}

	userModel := *params.ToModel()

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = model.CreateUser(db, &userModel)
	if err != nil {
		logrus.Error(err)
		return err
	}

	controller.EncodeJSONResponse(w, userviews.SingleFrom(userModel))
	w.WriteHeader(http.StatusCreated)
	return nil
}

func HandleBatchCreate(w http.ResponseWriter, r *http.Request) error {
	err := auth.CheckPermissions(r, userpermissiongroups.Create())
	if err != nil {
		logrus.Error(err)
		return apierrors.ClientForbiddenError{
			Message: fmt.Sprintf("Error batch creating users: %v", err),
		}
	}

	var usersparams userparams.BatchCreate
	if err := json.NewDecoder(r.Body).Decode(&usersparams); err != nil {
		e, ok := err.(*json.UnmarshalTypeError)
		if !ok {
			logrus.Error(err)
			return apierrors.ClientBadRequestError{
				Message: fmt.Sprintf("Bad JSON parsing: %v", err),
			}
		}

		// TODO: Investigate if we should use errors.Wrap instead
		return apierrors.ClientUnprocessableEntityError{
			Message: fmt.Sprintf("Invalid JSON format: %s should be a %s.", e.Field, e.Type),
		}
	}

	for _, userparams := range usersparams.Users {
		err := userparams.Validate()
		if err != nil {
			logrus.Error(err)
			return apierrors.ClientUnprocessableEntityError{
				Message: fmt.Sprintf("JSON validation failed: %v", err),
			}
		}
	}

	userModels := make([]*model.User, len(usersparams.Users))
	for i, userparams := range usersparams.Users {
		userModels[i] = userparams.ToModel()
	}

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	numCreated, err := model.CreateUsers(db, &userModels)
	if err != nil {
		logrus.Error(err)
		return err
	}

	controller.EncodeJSONResponse(w, userviews.BatchCreateFrom(userModels, numCreated))
	w.WriteHeader(http.StatusCreated)
	return nil
}
