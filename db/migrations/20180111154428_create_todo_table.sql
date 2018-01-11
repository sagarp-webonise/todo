
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

create table todo(
  id  SERIAL PRIMARY KEY,
  title VARCHAR(100) NOT NULL UNIQUE,
  done boolean
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

drop table todo;
