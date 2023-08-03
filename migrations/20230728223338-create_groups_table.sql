-- +migrate Up

CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    name TEXT,
    created_at created_at_type,
    deleted_at deleted_at_type,
    updated_at updated_at_type
);

-- +migrate Down

DROP TABLE IF EXISTS groups;
