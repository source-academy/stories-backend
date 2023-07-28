-- +migrate Up

CREATE TYPE role AS ENUM('member', 'moderator', 'admin')

CREATE TABLE IF NOT EXISTS usergroups (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    group_id INT REFERENCES groups(id),
    role role NOT NULL DEFAULT 'member',
    created_at created_at_type,
    deleted_at  deleted_at_type,
    updated_at updated_at_type
);

-- +migrate Down

DROP TABLE IF EXISTS usergroups;