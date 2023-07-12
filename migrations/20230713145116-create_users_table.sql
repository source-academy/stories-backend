-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    github_username TEXT,
    github_ID INTEGER,
    created_at created_at_type,
    deleted_at  deleted_at_type,
    updated_at updated_at_type
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
