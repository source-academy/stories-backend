package userpermissions

type Permission string

const (
	CanCreateUsers Permission = "can_create_users"
	CanReadUsers   Permission = "can_read_users"
	CanUpdateUsers Permission = "can_update_users"
	CanDeleteUsers Permission = "can_delete_users"

	CanCreateGroups Permission = "can_create_groups"
	CanReadGroups   Permission = "can_read_groups"
	CanUpdateGroups Permission = "can_update_groups"
	CanDeleteGroups Permission = "can_delete_groups"

	CanCreateStories  Permission = "can_create_stories"
	CanReadStories    Permission = "can_read_stories"
	CanUpdateStories  Permission = "can_update_stories"
	CanDeleteStories  Permission = "can_delete_stories"
	CanPublishStories Permission = "can_publish_stories"
)
