
-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    github_username TEXT,
    github_ID INTEGER,
    UNIQUE (user_id)
);
-- +migrate Down
