package usergroups

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/internal/auth"
	"github.com/source-academy/stories-backend/internal/database"
	apierrors "github.com/source-academy/stories-backend/internal/errors"
	userpermissiongroups "github.com/source-academy/stories-backend/internal/permissiongroups/users"
	"github.com/source-academy/stories-backend/model"
	usergroupparams "github.com/source-academy/stories-backend/params/usergroups"
	userviews "github.com/source-academy/stories-backend/view/users"
)

func HandleUpdateRole(w http.ResponseWriter, r *http.Request) error {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return apierrors.ClientBadRequestError{
			Message: fmt.Sprintf("Invalid userID: %v", err),
		}
	}

	err = auth.CheckPermissions(r, userpermissiongroups.Update(uint(userID)))
	if err != nil {
		logrus.Error(err)
		return apierrors.ClientForbiddenError{
			Message: fmt.Sprintf("Error batch creating users: %v", err),
		}
	}

	var params usergroupparams.UpdateRole
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

	// TODO: add validation for updating user?
	// err = params.Validate()
	// if err != nil {
	// 	logrus.Error(err)
	// 	return apierrors.ClientUnprocessableEntityError{
	// 		Message: fmt.Sprintf("JSON validation failed: %v", err),
	// 	}
	// }

	userGroupModel := *params.ToModel()

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	userGroup, err := model.UpdateUserGroupByUserID(db, uint(userID), &userGroupModel)
	if err != nil {
		logrus.Error(err)
		return err
	}

	user, err := model.GetUserByID(db, userID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	controller.EncodeJSONResponse(w, userviews.SummaryFrom(user, userGroup))
	return nil
}
