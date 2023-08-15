package stories

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/internal/database"
	apierrors "github.com/source-academy/stories-backend/internal/errors"
	"github.com/source-academy/stories-backend/internal/usergroups"
	"github.com/source-academy/stories-backend/model"
	storyparams "github.com/source-academy/stories-backend/params/stories"
	storyviews "github.com/source-academy/stories-backend/view/stories"
)

func HandleList(w http.ResponseWriter, r *http.Request) error {
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

func HandleRead(w http.ResponseWriter, r *http.Request) error {
	storyIDStr := chi.URLParam(r, "storyID")
	storyID, err := strconv.Atoi(storyIDStr)
	if err != nil {
		return apierrors.ClientBadRequestError{
			Message: fmt.Sprintf("Invalid storyID: %v", err),
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
	controller.EncodeJSONResponse(w, storyviews.SingleFrom(story))
	return nil
}

func HandleCreate(w http.ResponseWriter, r *http.Request) error {
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

	err := params.Validate()
	if err != nil {
		logrus.Error(err)
		return apierrors.ClientUnprocessableEntityError{
			Message: fmt.Sprintf("JSON validation failed: %v", err),
		}
	}

	storyModel := *params.ToModel()

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
