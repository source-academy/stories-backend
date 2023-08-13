package storypermissiongroups

import (
	"github.com/source-academy/stories-backend/internal/permissions"
	userpermissions "github.com/source-academy/stories-backend/internal/permissions/users"
)

func List() permissions.PermissionGroup {
	return userpermissions.
		GetRolePermission(userpermissions.CanReadStories)
}

func Create() permissions.PermissionGroup {
	return userpermissions.
		GetRolePermission(userpermissions.CanCreateStories)
}

func Read() permissions.PermissionGroup {
	return userpermissions.
		GetRolePermission(userpermissions.CanReadStories)
}

func Update(storyID uint) permissions.PermissionGroup {
	return permissions.AnyOf{
		Groups: []permissions.PermissionGroup{
			userpermissions.
				GetRolePermission(userpermissions.CanUpdateStories),
			IsAuthorOf{StoryID: storyID},
		},
	}
}

func Delete(storyID uint) permissions.PermissionGroup {
	return permissions.AnyOf{
		Groups: []permissions.PermissionGroup{
			userpermissions.
				GetRolePermission(userpermissions.CanDeleteStories),
			IsAuthorOf{StoryID: storyID},
		},
	}
}
