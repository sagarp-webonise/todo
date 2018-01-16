
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

update todo set done=false;
alter table todo alter COLUMN  done set not null;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
