-- +goose Up
-- +goose StatementBegin
CREATE TABLE authors(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL
);
CREATE TABLE books(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    author_id INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE authors;
DROP TABLE books;
-- +goose StatementEnd
