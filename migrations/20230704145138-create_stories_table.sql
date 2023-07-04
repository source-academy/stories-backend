-- +migrate Up

CREATE TABLE IF NOT EXISTS stories (
    story_id SERIAL PRIMARY KEY,
    user_id INT,
    story_content TEXT,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    UNIQUE (story_id)
);

-- +migrate Down