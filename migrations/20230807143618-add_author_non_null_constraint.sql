-- +migrate Up
ALTER TABLE stories
    ALTER COLUMN author_id SET NOT NULL;

-- -- +migrate Down
ALTER TABLE stories
    ALTER COLUMN author_id DROP NOT NULL;