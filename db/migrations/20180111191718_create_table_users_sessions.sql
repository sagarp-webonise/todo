-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE users_sessions (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id),
  session_id VARCHAR(255) NOT NULL,
  modified TIMESTAMP,
  created TIMESTAMP
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE users_sessions;
