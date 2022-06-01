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
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    PRIMARY KEY(code)
);

CREATE OR REPLACE FUNCTION ports_trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.modified_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON ports
FOR EACH ROW
EXECUTE PROCEDURE ports_trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS ports;
-- +goose StatementEnd
