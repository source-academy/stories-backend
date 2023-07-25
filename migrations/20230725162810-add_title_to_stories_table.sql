-- +migrate Up

ALTER TABLE stories
ADD COLUMN title VARCHAR(235);

-- +migrate Down

ALTER TABLE stories
DROP COLUMN title;
