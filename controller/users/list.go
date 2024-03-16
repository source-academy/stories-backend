package users

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/internal/auth"
	"github.com/source-academy/stories-backend/internal/database"
	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
	apierrors "github.com/source-academy/stories-backend/internal/errors"
	userpermissiongroups "github.com/source-academy/stories-backend/internal/permissiongroups/users"
	"github.com/source-academy/stories-backend/internal/usergroups"
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

	groupId, err := usergroups.GetGroupIDFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	roles := make([]groupenums.Role, len(users))
	for i, user := range users {
		roles[i], err = model.GetUserRoleByID(db, user.ID, *groupId)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	controller.EncodeJSONResponse(w, userviews.ListFrom(users, roles))
	return nil
}
