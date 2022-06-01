-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS currencies (
    code VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    PRIMARY KEY(code)
);

CREATE TABLE IF NOT EXISTS countries (
    id SERIAL,
    name VARCHAR(255),
    iso_code VARCHAR(32),
    currency_code VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    PRIMARY KEY(id),
    CONSTRAINT idx_currency_code FOREIGN KEY(currency_code) REFERENCES currencies(code)
);

CREATE TABLE IF NOT EXISTS cities (
    id SERIAL,
    name VARCHAR(255),
    country_id int,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    PRIMARY KEY(id),
    CONSTRAINT idx_cities FOREIGN KEY(country_id) REFERENCES countries(id)
);

CREATE OR REPLACE FUNCTION currencies_trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.modified_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON currencies
FOR EACH ROW
EXECUTE PROCEDURE currencies_trigger_set_timestamp();

CREATE OR REPLACE FUNCTION countries_trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.modified_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON countries
FOR EACH ROW
EXECUTE PROCEDURE countries_trigger_set_timestamp();

CREATE OR REPLACE FUNCTION cities_trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.modified_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON cities
FOR EACH ROW
EXECUTE PROCEDURE cities_trigger_set_timestamp();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE countries DROP CONSTRAINT idx_currency_code;

ALTER TABLE cities DROP CONSTRAINT idx_cities;

DROP TABLE IF EXISTS cities;
DROP TABLE IF EXISTS countries;
DROP TABLE IF EXISTS currencies;
-- +goose StatementEnd
