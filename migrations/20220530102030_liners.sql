-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS liners (
    id SERIAL,
    name VARCHAR(255),
    code VARCHAR(255),
    type VARCHAR(32),
    logo VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    PRIMARY KEY(id)
);

CREATE OR REPLACE FUNCTION liners_trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.modified_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON liners
FOR EACH ROW
EXECUTE PROCEDURE liners_trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS liners;
-- +goose StatementEnd
