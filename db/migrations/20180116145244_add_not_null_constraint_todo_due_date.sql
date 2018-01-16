
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
update todo set due_date=now()+interval '48 hours';
alter table todo alter COLUMN  due_date set not null;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
