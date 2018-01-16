
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE  todo ALTER COLUMN due_date SET DEFAULT now();
UPDATE todo SET due_date = now();
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
