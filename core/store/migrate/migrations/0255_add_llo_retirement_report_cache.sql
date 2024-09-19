-- +goose Up
-- +goose StatementBegin

CREATE TABLE retirement_report_cache (
    config_digest BYTEA NOT NULL CHECK (OCTET_LENGTH(config_digest) = 32) PRIMARY KEY,
    attested_retirement_report BYTEA NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE retirement_report_cache;

-- +goose StatementEnd
