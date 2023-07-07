-- +migrate Up

CREATE TABLE IF NOT EXISTS stories (
    story_id SERIAL PRIMARY KEY,
    user_id INT,
    story_content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION update_stories_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger on the stories table to update the updated_at column
CREATE TRIGGER update_stories_updated_at_trigger
    BEFORE UPDATE ON stories
    FOR EACH ROW
    EXECUTE FUNCTION update_stories_updated_at_column();

-- +migrate StatementEnd

-- +migrate Down

DROP TRIGGER IF EXISTS update_stories_updated_at_trigger ON stories;
DROP FUNCTION IF EXISTS update_stories_updated_at_column();
DROP TABLE IF EXISTS stories;
