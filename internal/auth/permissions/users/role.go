package userpermissions

import (
	"net/http"
)

type RolePermission struct {
	Permission Permission
}

func (p RolePermission) IsAuthorized(r *http.Request) bool {
	// TODO: Implement. Blocked by request context missing
	// user and group info.
	return false
}
