-- +migrate Up

CREATE DOMAIN role_type ENUM('member', 'moderator', 'admin') NOT NULL DEFAULT 'member';

CREATE TABLE IF NOT EXISTS user_groups (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    group_id INT REFERENCES groups(id),
    role role_type,
    created_at created_at_type,
    deleted_at  deleted_at_type,
    updated_at updated_at_type
);

-- +migrate Down

DROP TABLE IF EXISTS user_groups;

DROP DOMAIN role_type;
