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
	apierrors "github.com/source-academy/stories-backend/internal/errors"
	"github.com/source-academy/stories-backend/model"
)

const (
	userGroupKey = "user_group_context"
)

func InjectUserGroupIntoContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		groupIDStr := chi.URLParam(r, "groupID")
		groupID, err := strconv.Atoi(groupIDStr)
		if err != nil {
			apierrors.ServeHTTP(w, r, apierrors.ClientBadRequestError{
				Message: fmt.Sprintf("Invalid groupID: %v", err),
			})
			return
		}

		// get user
		userID, err := auth.GetUserIDFrom(r)
		if err != nil {
			apierrors.ServeHTTP(w, r, apierrors.InternalServerError{
				Message: fmt.Sprintf("Failed to get user: %s\n", err),
			})
			return
		}

		// Get DB instance
		db, err := database.GetDBFrom(r)
		if err != nil {
			logrus.Error(err)
			apierrors.ServeHTTP(w, r, apierrors.InternalServerError{
				Message: fmt.Sprintf("Failed to get DB instance: %s\n", err),
			})
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
			apierrors.ServeHTTP(w, r, apierrors.ClientForbiddenError{
				Message: fmt.Sprintf("User not in group: %s\n", err),
			})
			return
		}

		// Create a new session from the database for each request
		ctx := context.WithValue(r.Context(), userGroupKey, dbUserGroup)
		// Inject the new context into the request
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserGroupFrom(r *http.Request) (*model.UserGroup, error) {
	userGroup, ok := r.Context().Value(userGroupKey).(*model.UserGroup)
	if !ok {
		return nil, errors.New("Could not get database from request context")
	}
	return userGroup, nil
}
