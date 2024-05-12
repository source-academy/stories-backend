package userpermissiongroups

import (
	"github.com/source-academy/stories-backend/internal/permissions"
	userpermissions "github.com/source-academy/stories-backend/internal/permissions/users"
)

func List() permissions.PermissionGroup {
	return userpermissions.
		GetRolePermission(userpermissions.CanReadUsers)
}

func Create() permissions.PermissionGroup {
	return userpermissions.
		GetRolePermission(userpermissions.CanCreateUsers)
}

func Read(userID uint) permissions.PermissionGroup {
	return permissions.AnyOf{
		Groups: []permissions.PermissionGroup{
			userpermissions.
				GetRolePermission(userpermissions.CanReadUsers),
			IsSelf{UserID: userID},
		},
	}
}

func Update(userID uint) permissions.PermissionGroup {
	return permissions.AnyOf{
		Groups: []permissions.PermissionGroup{
			userpermissions.
				GetRolePermission(userpermissions.CanUpdateUsers),
		},
	}
}

func Delete() permissions.PermissionGroup {
	return userpermissions.
		GetRolePermission(userpermissions.CanDeleteUsers)
}
