-- +migrate Up

ALTER TABLE users
    ADD COLUMN full_name TEXT NOT NULL DEFAULT '';

-- +migrate Down

ALTER TABLE users
    DROP COLUMN full_name;
