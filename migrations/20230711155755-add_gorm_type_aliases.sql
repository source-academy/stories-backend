-- Creates the GORM type aliases for the domains used in the models.
-- Now, we can just create new tables like
-- CREATE TABLE users (
--   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--   created_at created_at_type,
--   updated_at updated_at_type,
--   deleted_at deleted_at_type
-- );
-- instead of
-- CREATE TABLE users (
--   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--   created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
--   updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
--   deleted_at TIMESTAMPTZ DEFAULT NULL
-- );

-- +migrate Up

CREATE DOMAIN created_at_type TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP;
CREATE DOMAIN updated_at_type TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP;
CREATE DOMAIN deleted_at_type TIMESTAMPTZ DEFAULT NULL;

-- +migrate Down

DROP DOMAIN created_at_type;
DROP DOMAIN updated_at_type;
DROP DOMAIN deleted_at_type;
