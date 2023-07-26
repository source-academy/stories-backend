-- +migrate Up

ALTER TABLE stories
    ADD COLUMN pin_order INT;

-- +migrate Down

ALTER TABLE stories
    DROP COLUMN pin_order;
