-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    github_username TEXT,
    github_ID INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE (user_id)
);

-- +migrate Down
