-- +goose Up
-- +goose StatementBegin
CREATE TABLE branches (
    id INTEGER,
    type INTEGER,
    forumId INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE branches;
-- +goose StatementEnd
