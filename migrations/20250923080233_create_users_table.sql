-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id SERIAL not null primary key,
  first_name varchar,
  last_name varchar,
  bio varchar
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
