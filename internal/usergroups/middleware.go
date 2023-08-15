package usergroups

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/internal/auth"
	"github.com/source-academy/stories-backend/internal/database"
	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
	apierrors "github.com/source-academy/stories-backend/internal/errors"
	"github.com/source-academy/stories-backend/model"
)

type contextKey string

const (
	groupKey contextKey = "group_context"
	roleKey  contextKey = "role_context"
)

func InjectUserGroupIntoContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		groupIDStr := chi.URLParam(r, "groupID")
		groupID, err := strconv.Atoi(groupIDStr)
		if err != nil {
			logrus.Error(err)
			apierrors.ServeHTTP(w, r, apierrors.ClientNotFoundError{
				Message: "Route not found.",
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
			UserID:  uint(*userID),
			GroupID: uint(groupID),
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

		groupContext := context.WithValue(r.Context(), groupKey, &dbUserGroup.GroupID)
		ctx := context.WithValue(groupContext, roleKey, dbUserGroup.Role)
		// Inject the new context into the request
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetGroupIDFrom(r *http.Request) (*uint, error) {
	groupID, ok := r.Context().Value(groupKey).(*uint)
	if !ok {
		return nil, errors.New("Could not get groupID from request context")
	}
	return groupID, nil
}

func GetRoleFrom(r *http.Request) (*groupenums.Role, error) {
	role, ok := r.Context().Value(roleKey).(groupenums.Role)
	if !ok {
		return nil, errors.New("Could not get role from request context")
	}
	return &role, nil
}
