-- +migrate Up

ALTER TABLE stories
    ADD COLUMN group_id INT REFERENCES groups(id);

-- +migrate Down

ALTER TABLE stories
    DROP COLUMN group_id;
