-- +migrate Up

CREATE UNIQUE INDEX idx_unique_username_provider ON users (username, login_provider);

-- +migrate Down

DROP INDEX idx_unique_username_provider;
