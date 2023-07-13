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

DROP TRIGGER IF EXISTS update_users_updated_at_trigger ON users;
DROP FUNCTION IF EXISTS update_users_updated_at_column();
DROP TABLE IF EXISTS users;
