-- +migrate Up

CREATE TABLE IF NOT EXISTS stories (
    story_id SERIAL PRIMARY KEY,
    user_id INT,
    story_content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    UNIQUE (story_id)
);

-- +migrate Down

DROP TABLE IF EXISTS stories;
