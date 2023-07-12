-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    github_username TEXT,
    github_ID INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION update_users_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    IF (NEW.github_username IS DISTINCT FROM OLD.github_username) THEN
        NEW.updated_at = CURRENT_TIMESTAMP;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_users_updated_at_trigger
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_users_updated_at_column();

-- +migrate StatementEnd

-- +migrate Down

DROP TRIGGER IF EXISTS update_users_updated_at_trigger ON users;
DROP FUNCTION IF EXISTS update_users_updated_at_column();
DROP TABLE IF EXISTS users;
