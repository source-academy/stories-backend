package userpermissions

import (
	"errors"

	groupenums "github.com/source-academy/stories-backend/internal/enums/groups"
)

// Gets the RolePermission from a Permission. To be called
// by external packages.
// TODO: Investigate possibility of defining roles on a
// per-user/group basis in the DB.
func GetRolePermission(p Permission) *RolePermission {
	switch p {
	case
		// Permissions for all users
		CanCreateStories,
		CanReadStories:
		return &RolePermission{
			Permission: p,
			Role:       groupenums.RoleStandard,
		}
	case
		// Additional permissions for moderators and administrators
		CanUpdateStories,
		CanDeleteStories:
		return &RolePermission{
			Permission: p,
			Role:       groupenums.RoleModerator,
		}
	case
		// Additional permissions for administrators only
		CanCreateUsers,
		CanReadUsers,
		CanUpdateUsers,
		CanDeleteUsers,
		CanCreateGroups,
		CanReadGroups,
		CanUpdateGroups,
		CanDeleteGroups:
		return &RolePermission{
			Permission: p,
			Role:       groupenums.RoleAdmin,
		}
	}
	// Illegal path - all permissions should have been handled above
	panic(errors.New("Illegal path - all permissions should have been handled above"))
}
