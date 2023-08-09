package userpermissions

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/internal/auth"
	"github.com/source-academy/stories-backend/internal/database"
	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
	apierrors "github.com/source-academy/stories-backend/internal/errors"
	"github.com/source-academy/stories-backend/model"
)

type contextKey string

const (
	GroupKey contextKey = "group_context"
	RoleKey  contextKey = "role_context"
)

func InjectUserGroupIntoContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		groupIDStr := chi.URLParam(r, "groupID")
		groupID, err := strconv.Atoi(groupIDStr)
		if err != nil {
			logrus.Error(err)
			apierrors.ServeHTTP(w, r, apierrors.ClientUnprocessableEntityError{
				Message: fmt.Sprintf("Invalid groupID: %v", err),
			})
			return
		}

		// Get user, this is called after the user ID injection
		userID, err := auth.GetUserIDFrom(r)
		if err != nil {
			logrus.Error(err)
			apierrors.ServeHTTP(w, r, err)
			return
		}

		// Get DB instance, this is called after the DB middleware
		db, err := database.GetDBFrom(r)
		if err != nil {
			logrus.Error(err)
			apierrors.ServeHTTP(w, r, err)
			return
		}

		// Get user_group
		userGroup := model.UserGroup{
			UserID:  *userID,
			GroupID: groupID,
		}
		var dbUserGroup model.UserGroup
		err = db.Where(&userGroup).First(&dbUserGroup).Error
		if err != nil {
			logrus.Error(err)
			apierrors.ServeHTTP(w, r, apierrors.ClientForbiddenError{
				Message: database.HandleDBError(err, "user_group").Error(),
			})
			return
		}

		group_context := context.WithValue(r.Context(), GroupKey, dbUserGroup.GroupID)
		ctx := context.WithValue(group_context, RoleKey, dbUserGroup.Role)
		// Inject the new context into the request
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetGroupIDFrom(r *http.Request) (*int, error) {
	groupID, ok := r.Context().Value(GroupKey).(*int)
	if !ok {
		return nil, errors.New("Could not get groupID from request context")
	}
	return groupID, nil
}

func GetRoleDFrom(r *http.Request) (*groupenums.Role, error) {
	role, ok := r.Context().Value(RoleKey).(*groupenums.Role)
	if !ok {
		return nil, errors.New("Could not get role from request context")
	}
	return role, nil
}
