package auth

import (
	"net/http"

	"github.com/source-academy/stories-backend/internal/auth/permissions"
)

func CheckPermissions(r *http.Request, requestedActionPermissions ...permissions.PermissionGroup) (bool, error) {
	// TODO: check permissions
	// TODO: perhaps utilize userpermissions.GetFromRole
	return false, nil
}
