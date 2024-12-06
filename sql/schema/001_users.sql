-- +goose Up
CREATE TABLE users (
    id uuid unique primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text unique not null
);
-- +goose Down
DROP TABLE users;