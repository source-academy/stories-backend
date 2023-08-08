package auth

import (
	"errors"
	"net/http"

	"github.com/source-academy/stories-backend/internal/permissions"
)

func CheckPermissions(r *http.Request, requestedActionPermissions ...permissions.PermissionGroup) error {
	requiredPermissions := permissions.AllOf{
		Groups: requestedActionPermissions,
	}
	isAuthorized := requiredPermissions.IsAuthorized(r)
	if !isAuthorized {
		return errors.New("You are not authorized to perform this action.")
	}
	return nil
}
