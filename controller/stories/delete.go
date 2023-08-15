package stories

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/controller"
	"github.com/source-academy/stories-backend/internal/database"
	apierrors "github.com/source-academy/stories-backend/internal/errors"
	"github.com/source-academy/stories-backend/model"
	storyviews "github.com/source-academy/stories-backend/view/stories"
)

func HandleDelete(w http.ResponseWriter, r *http.Request) error {
	storyIDStr := chi.URLParam(r, "storyID")
	storyID, err := strconv.Atoi(storyIDStr)
	if err != nil {
		return apierrors.ClientBadRequestError{
			Message: fmt.Sprintf("Invalid storyID: %v", err),
		}
	}

	// TODO: Prevents cross-tenant story viewing
	//       when user is a member of multiple stories groups.
	//       Not implemented yet as deletion is protected with more
	//       stringent permissions checks.
	//       Will await refactor to minimise wasted effort

	// Get DB instance
	db, err := database.GetDBFrom(r)
	if err != nil {
		logrus.Error(err)
		return err
	}

	story, err := model.DeleteStory(db, storyID)
	if err != nil {
		logrus.Error(err)
		return err
	}
	controller.EncodeJSONResponse(w, storyviews.SingleFrom(story))
	return nil
}
