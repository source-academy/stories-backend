-- +migrate Up

CREATE TABLE IF NOT EXISTS stories (
    story_id SERIAL PRIMARY KEY,
    author_id INT REFERENCES users(user_id),
    story_content TEXT,
    created_at created_at_type,
    deleted_at  deleted_at_type,
    updated_at updated_at_type
);

-- +migrate Down

DROP TRIGGER IF EXISTS update_stories_updated_at_trigger ON stories;
DROP FUNCTION IF EXISTS update_stories_updated_at_column();
DROP TABLE IF EXISTS stories;
