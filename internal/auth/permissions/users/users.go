package userpermissions

import (
	"github.com/source-academy/stories-backend/internal/auth/permissions"
)

// TODO: Replace with more maintainable permissions matrix,
// perhaps defined on a per-user/group basis in the DB.
// TODO: Accept role as parameter once role enum is merged.
// func GetFromRole(role groupenums.Role) []permissions.PermissionGroup {
func GetFromRole() []permissions.PermissionGroup {
	permissionsList := &[]Permission{} // FIXME: Hack to prevent errors. Remove when role enums are merged.
	// var permissionsList *[]Permission

	// Get permissions for all users
	standardPermissions := []Permission{
		CanCreateStories,
		CanReadStories,
	}
	// if role == groupenums.RoleStandard {
	// 	permissionsList = &standardPermissions
	// }

	// Get additional permissions for moderators and administrators
	moderatorPermissions := []Permission{
		CanUpdateStories,
		CanDeleteStories,
	}
	moderatorPermissions = append(moderatorPermissions, standardPermissions...)
	// if role == groupenums.RoleModerator {
	// 	permissionsList = &moderatorPermissions
	// }

	// Get additional permissions for administrators only
	adminPermissions := []Permission{
		CanCreateUsers,
		CanReadUsers,
		CanUpdateUsers,
		CanDeleteUsers,
		CanCreateGroups,
		CanReadGroups,
		CanUpdateGroups,
		CanDeleteGroups,
	}
	adminPermissions = append(adminPermissions, moderatorPermissions...)
	// if role == groupenums.RoleAdmin {
	// 	permissionsList = &adminPermissions
	// }

	// Return the permissions group

	_ = adminPermissions // FIXME: Hack to prevent lint errors. Remove when role enums are merged.

	if permissionsList == nil {
		// Illegal path - something must have went wrong
		panic("Illegal path")
	}

	group := make([]permissions.PermissionGroup, 0, len(standardPermissions))
	for i, p := range *permissionsList {
		group[i] = RolePermission{Permission: p}
	}
	return group
}
