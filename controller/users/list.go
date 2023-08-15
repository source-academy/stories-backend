package users

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/internal/auth"
	"github.com/source-academy/stories-backend/internal/database"
	apierrors "github.com/source-academy/stories-backend/internal/errors"
	userpermissiongroups "github.com/source-academy/stories-backend/internal/permissiongroups/users"
	"github.com/source-academy/stories-backend/model"
	userviews "github.com/source-academy/stories-backend/view/users"
)

func HandleList(w http.ResponseWriter, r *http.Request) error {
	err := auth.CheckPermissions(r, userpermissiongroups.List())
	if err != nil {
		logrus.Error(err)
		return apierrors.ClientForbiddenError{
			Message: fmt.Sprintf("Error listing users: %v", err),
		}
	}

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
