-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS hs_codes (
    code VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    description VARCHAR(255),
    parent_code VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    modified_at TIMESTAMP NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    PRIMARY KEY(code)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS hs_codes;
-- +goose StatementEnd
