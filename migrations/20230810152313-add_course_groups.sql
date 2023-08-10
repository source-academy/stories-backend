-- +migrate Up

CREATE TABLE IF NOT EXISTS course_groups (
    id SERIAL,
    course_id INT PRIMARY KEY,
    group_id INT REFERENCES groups(id) NOT NULL,
    created_at created_at_type,
    deleted_at  deleted_at_type,
    updated_at updated_at_type
);

-- +migrate Down

DROP TABLE IF EXISTS user_groups;

