-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS organizations(
    id UUID DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    address TEXT NOT NULL UNIQUE,
    email_address VARCHAR(255) NOT NULL UNIQUE,
    phone_number VARCHAR(20) NOT NULL UNIQUE,
    owner_id UUID NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS api_keys(
    id UUID DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    test_secret_key TEXT NOT NULL UNIQUE,
    test_public_key TEXT NOT NULL UNIQUE,
    live_secret_key TEXT NOT NULL UNIQUE,
    live_public_key TEXT NOT NULL UNIQUE,
    owner_id UUID NOT NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS organizations;
-- +goose StatementEnd
