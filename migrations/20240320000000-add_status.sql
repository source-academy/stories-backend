-- +migrate Up

ALTER TABLE stories
    ADD COLUMN status INT;

-- +migrate Down

ALTER TABLE stories
    DROP COLUMN status;