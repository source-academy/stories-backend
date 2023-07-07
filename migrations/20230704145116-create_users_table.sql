-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    github_username TEXT,
    github_ID INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION update_users_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger on the users table to update the updated_at column
CREATE TRIGGER update_users_updated_at_trigger
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_users_updated_at_column();

-- +migrate StatementEnd

-- +migrate Down

DROP TRIGGER IF EXISTS update_users_updated_at_trigger ON users;
DROP FUNCTION IF EXISTS update_users_updated_at_column();
DROP TABLE IF EXISTS users;
