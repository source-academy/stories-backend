-- +migrate Up

CREATE TABLE IF NOT EXISTS course_groups (
    id SERIAL PRIMARY KEY,
    course_id INT NOT NULL,
    group_id INT REFERENCES groups(id) NOT NULL,
    created_at created_at_type,
    deleted_at deleted_at_type,
    updated_at updated_at_type
);

CREATE UNIQUE INDEX idx_unique_course ON course_groups (course_id);

-- +migrate Down

DROP INDEX idx_unique_course;

DROP TABLE IF EXISTS course_groups;

