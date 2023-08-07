-- +migrate Up
ALTER TABLE user_groups
    ALTER COLUMN user_id SET NOT NULL,
    ALTER COLUMN group_id SET NOT NULL;

-- +migrate Down
ALTER TABLE user_groups
    ALTER COLUMN user_id DROP NOT NULL,
    ALTER COLUMN group_id DROP NOT NULL;
