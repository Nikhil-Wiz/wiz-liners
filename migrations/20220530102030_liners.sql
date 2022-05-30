-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS liners (
    id int NOT NULL,
    name VARCHAR(255),
    code VARCHAR(255),
    type VARCHAR(32),
    logo VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    modified_at TIMESTAMP NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS liners;
-- +goose StatementEnd
