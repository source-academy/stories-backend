-- Makes the users table agnostic to the login provider.

-- +migrate Up

CREATE DOMAIN login_provider_type INTEGER NOT NULL;

ALTER TABLE users
    ADD COLUMN username TEXT NOT NULL,
    -- Do not enforce the non null constraint yet.
    ADD COLUMN login_provider INTEGER;

-- Begin data migration
-- NOTE: We ignore the data migration here.
-- End data migration

ALTER TABLE users
    DROP COLUMN github_username,
    DROP COLUMN github_id;

ALTER TABLE users
    -- Enforce the non null constraint.
    ALTER COLUMN login_provider TYPE login_provider_type;

-- +migrate Down

ALTER TABLE users
    DROP COLUMN username,
    DROP COLUMN login_provider;

-- NOTE: We ignore the data migration here.

ALTER TABLE users
    ADD COLUMN github_username TEXT,
    ADD COLUMN github_id INTEGER;

DROP DOMAIN login_provider_type;
