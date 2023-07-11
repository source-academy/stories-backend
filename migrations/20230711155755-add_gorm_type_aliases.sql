-- Creates the GORM type aliases for the domains used in the models.

-- +migrate Up

CREATE DOMAIN created_at_type TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP;
CREATE DOMAIN updated_at_type TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP;
CREATE DOMAIN deleted_at_type TIMESTAMPTZ DEFAULT NULL;

-- +migrate Down

DROP DOMAIN created_at_type;
DROP DOMAIN updated_at_type;
DROP DOMAIN deleted_at_type;
