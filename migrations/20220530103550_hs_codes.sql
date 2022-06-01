-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS hs_codes (
    code VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    description VARCHAR(255),
    parent_code VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    PRIMARY KEY(code)
);

CREATE OR REPLACE FUNCTION hscodes_trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.modified_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON hs_codes
FOR EACH ROW
EXECUTE PROCEDURE hscodes_trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS hs_codes;
-- +goose StatementEnd
