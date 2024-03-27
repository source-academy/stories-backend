-- +migrate Up

ALTER TABLE stories
    ADD COLUMN status_message TEXT;

-- +migrate Down

ALTER TABLE stories
    DROP COLUMN status_message;