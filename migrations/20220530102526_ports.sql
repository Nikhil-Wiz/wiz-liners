-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ports (
    code VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    type VARCHAR(32),
    city_id int,
    state VARCHAR(255),
    latitude DECIMAL(5,3),
    longitude DECIMAL(5,3),
    created_at TIMESTAMP NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    modified_at TIMESTAMP NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    PRIMARY KEY(code)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS ports;
-- +goose StatementEnd
