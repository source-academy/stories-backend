package auth

import (
	"net/http"

	"github.com/source-academy/stories-backend/internal/permissions"
)

func CheckPermissions(r *http.Request, requestedActionPermissions ...permissions.PermissionGroup) (bool, error) {
	requiredPermissions := permissions.AllOf{
		Groups: requestedActionPermissions,
	}
	return requiredPermissions.IsAuthorized(r), nil
}
