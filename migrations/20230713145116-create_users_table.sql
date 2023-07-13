-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    github_username TEXT,
    github_ID INTEGER,
    created_at created_at_type,
    deleted_at  deleted_at_type,
    updated_at updated_at_type
);

-- +migrate Down

DROP TABLE IF EXISTS users;
