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
	"github.com/source-academy/stories-backend/model"
	storyparams "github.com/source-academy/stories-backend/params/stories"
	storyviews "github.com/source-academy/stories-backend/view/stories"
)

func HandleUpdate(w http.ResponseWriter, r *http.Request) error {
	storyIDStr := chi.URLParam(r, "storyID")
	storyID, err := strconv.Atoi(storyIDStr)
	if err != nil {
		return apierrors.ClientBadRequestError{
			Message: fmt.Sprintf("Invalid storyID: %v", err),
		}
	}

	err = auth.CheckPermissions(r, storypermissiongroups.Update(uint(storyID)))
	if err != nil {
		logrus.Error(err)
		return apierrors.ClientForbiddenError{
			Message: fmt.Sprintf("Error updating story: %v", err),
		}
	}

	// Extra params won't do anything, e.g. authorID can't be changed.
	// TODO: Error on extra params?
	var params storyparams.Update
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

	storyModel := *params.ToModel()

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// TODO: Prevents cross-tenant story viewing
	//       when user is a member of multiple stories groups.
	//       Not implemented yet as deletion is protected with more
	//       stringent permissions checks.
	//       Will await refactor to minimise wasted effort

	err = model.UpdateStory(db, storyID, &storyModel)
	if err != nil {
		logrus.Error(err)
		return err
	}

	controller.EncodeJSONResponse(w, storyviews.SingleFrom(storyModel))
	return nil
}
