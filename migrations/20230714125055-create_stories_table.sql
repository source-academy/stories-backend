-- +migrate Up

CREATE TABLE IF NOT EXISTS stories (
    id SERIAL PRIMARY KEY,
    author_id INT UNSIGNED REFERENCES users(user_id),
    story_content TEXT,
    created_at created_at_type,
    deleted_at  deleted_at_type,
    updated_at updated_at_type
);

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION update_stories_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    IF (NEW.story_content IS DISTINCT FROM OLD.story_content) THEN
        NEW.updated_at = CURRENT_TIMESTAMP;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_stories_updated_at_trigger
    BEFORE UPDATE ON stories
    FOR EACH ROW
    EXECUTE FUNCTION update_stories_updated_at_column();

-- +migrate StatementEnd

-- +migrate Down

DROP TRIGGER IF EXISTS update_stories_updated_at_trigger ON stories;
DROP FUNCTION IF EXISTS update_stories_updated_at_column();
DROP TABLE IF EXISTS stories;
