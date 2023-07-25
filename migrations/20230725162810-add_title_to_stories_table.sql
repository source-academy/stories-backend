-- +migrate Up

ALTER TABLE stories
ADD COLUMN title TEXT;

-- +migrate Down

ALTER TABLE stories
DROP COLUMN title;
