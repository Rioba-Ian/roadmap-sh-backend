-- +goose Up
-- +goose StatementBegin
create extension if not exists citext;
create extension if not exists "uuid-ossp";

CREATE TABLE users (
  id uuid primary key default gen_random_uuid(),
  first_name varchar(255) not null,
  last_name varchar(255),
  email citext not null unique,
  password_hash varchar(255),
  token text,
  refresh_token text,
  created_at timestamp(0) not null default (now() at time zone 'utc'),
  updated_at timestamp(0) not null default (now() at time zone 'utc')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
