-- +migrate Up

CREATE TABLE IF NOT EXISTS stories (
    id SERIAL PRIMARY KEY,
    author_id INT REFERENCES users(id),
    content TEXT,
    created_at created_at_type,
    deleted_at  deleted_at_type,
    updated_at updated_at_type
);

-- +migrate Down

DROP TABLE IF EXISTS stories;
