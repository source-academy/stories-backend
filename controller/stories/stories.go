package stories

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
	storypermissiongroups "github.com/source-academy/stories-backend/internal/permissiongroups/stories"
	"github.com/source-academy/stories-backend/internal/usergroups"
	"github.com/source-academy/stories-backend/model"
	storyparams "github.com/source-academy/stories-backend/params/stories"
	storyviews "github.com/source-academy/stories-backend/view/stories"
)

func HandleList(w http.ResponseWriter, r *http.Request) error {
	err := auth.CheckPermissions(r, storypermissiongroups.List())
	if err != nil {
		logrus.Error(err)
		return apierrors.ClientForbiddenError{
			Message: fmt.Sprintf("Error listing stories: %v", err),
		}
	}

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// Get group id from context
	groupID, err := usergroups.GetGroupIDFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	stories, err := model.GetAllStoriesInGroup(db, groupID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	controller.EncodeJSONResponse(w, storyviews.ListFrom(stories))
	return nil
}

func HandleListPublished(w http.ResponseWriter, r *http.Request) error {
	err := auth.CheckPermissions(r, storypermissiongroups.List())
	if err != nil {
		logrus.Error(err)
		return apierrors.ClientForbiddenError{
			Message: fmt.Sprintf("Error listing published stories: %v", err),
		}
	}

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// Get group id from context
	groupID, err := usergroups.GetGroupIDFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	stories, err := model.GetAllStoriesByStatus(db, groupID, model.Published)
	if err != nil {
		logrus.Error(err)
		return err
	}

	controller.EncodeJSONResponse(w, storyviews.ListFrom(stories))
	return nil
}

func HandleListPending(w http.ResponseWriter, r *http.Request) error {
	err := auth.CheckPermissions(r, storypermissiongroups.Moderate())
	if err != nil {
		logrus.Error(err)
		return apierrors.ClientForbiddenError{
			Message: fmt.Sprintf("Error listing stories to be reviewed: %v", err),
		}
	}

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// Get group id from context
	groupID, err := usergroups.GetGroupIDFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	stories, err := model.GetAllStoriesByStatus(db, groupID, model.Pending)
	if err != nil {
		logrus.Error(err)
		return err
	}

	controller.EncodeJSONResponse(w, storyviews.ListFrom(stories))
	return nil
}

func HandleListDraft(w http.ResponseWriter, r *http.Request) error {
	err := auth.CheckPermissions(r, storypermissiongroups.List())
	if err != nil {
		logrus.Error(err)
		return apierrors.ClientForbiddenError{
			Message: fmt.Sprintf("Error listing drafts: %v", err),
		}
	}
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

	// Get group id from context
	groupID, err := usergroups.GetGroupIDFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	stories, err := model.GetAllAuthorStoriesByStatus(db, groupID, userID, model.Draft)
	if err != nil {
		logrus.Error(err)
		return err
	}

	controller.EncodeJSONResponse(w, storyviews.ListFrom(stories))
	return nil
}

func HandleListRejected(w http.ResponseWriter, r *http.Request) error {
	err := auth.CheckPermissions(r, storypermissiongroups.List())
	if err != nil {
		logrus.Error(err)
		return apierrors.ClientForbiddenError{
			Message: fmt.Sprintf("Error listing rejected stories: %v", err),
		}
	}
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

	// Get group id from context
	groupID, err := usergroups.GetGroupIDFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	stories, err := model.GetAllAuthorStoriesByStatus(db, groupID, userID, model.Rejected)
	if err != nil {
		logrus.Error(err)
		return err
	}

	controller.EncodeJSONResponse(w, storyviews.ListFrom(stories))
	return nil
}

func HandleRead(w http.ResponseWriter, r *http.Request) error {
	storyIDStr := chi.URLParam(r, "storyID")
	storyID, err := strconv.Atoi(storyIDStr)
	if err != nil {
		return apierrors.ClientBadRequestError{
			Message: fmt.Sprintf("Invalid storyID: %v", err),
		}
	}

	err = auth.CheckPermissions(r, storypermissiongroups.Read())
	if err != nil {
		logrus.Error(err)
		return apierrors.ClientForbiddenError{
			Message: fmt.Sprintf("Error reading story: %v", err),
		}
	}

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	story, err := model.GetStoryByID(db, storyID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// TODO: Refactor
	// Prevents cross-tenant story viewing
	// when user is a member of multiple stories groups

	// Get group id from context
	groupID, err := usergroups.GetGroupIDFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// FIXME: Hacky; use nullable types!
	if *story.GroupID != *groupID {
		// Do not expose multitenancy implementation
		return apierrors.ClientNotFoundError{
			Message: fmt.Sprintf("Story with ID %d not found", storyID),
		}
	}

	controller.EncodeJSONResponse(w, storyviews.SingleFrom(story))
	return nil
}

func HandleCreate(w http.ResponseWriter, r *http.Request) error {
	err := auth.CheckPermissions(r, storypermissiongroups.Create())
	if err != nil {
		logrus.Error(err)
		return apierrors.ClientForbiddenError{
			Message: fmt.Sprintf("Error creating story: %v", err),
		}
	}

	var params storyparams.Create
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

	// Get group id from context
	groupID, err := usergroups.GetGroupIDFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	storyModel := *params.ToModel(groupID)

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = model.CreateStory(db, &storyModel)
	if err != nil {
		logrus.Error(err)
		return err
	}

	controller.EncodeJSONResponse(w, storyviews.SingleFrom(storyModel))
	w.WriteHeader(http.StatusCreated)
	return nil
}
