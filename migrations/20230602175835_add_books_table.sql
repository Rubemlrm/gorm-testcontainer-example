-- +goose Up
-- +goose StatementBegin
create table books
(
  id         bigserial primary key,
  title   VARCHAR(120) not null,
  author   VARCHAR(12) not null,
  created_at timestamp without time zone default (now() at time zone 'utc'),
  updated_at timestamp without time zone default (now() at time zone 'utc')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop table if exists books;
-- +goose StatementEnd
