-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id UUID DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    email_address VARCHAR(255) NOT NULL UNIQUE,
    phone_number VARCHAR(20) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

CREATE TYPE coa_type AS ENUM ('asset', 'liability', 'equity', 'revenue', 'expensis');

CREATE TABLE IF NOT EXISTS chart_of_accounts(
    id UUID DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    account_number VARCHAR(20) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    type coa_type NOT NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS chart_of_accounts;
DROP TYPE IF EXISTS coa_type;
-- +goose StatementEnd
