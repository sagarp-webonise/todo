-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users(
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE,
  first_name VARCHAR(255) NULL,
  last_name VARCHAR(255) NULL,
  email VARCHAR(255) NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  login_attempts INT NOT NULL DEFAULT 0,
  modified TIMESTAMP,
  created TIMESTAMP
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE users;
