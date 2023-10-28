-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE block_status AS ENUM ('open', 'close');
CREATE TYPE ledger_account_status AS ENUM ('pending', 'approved');
CREATE TYPE transaction_status AS ENUM ('pending', 'approved');
CREATE TYPE ledger_book AS ENUM ('general_journal', 'cash_receipt');
CREATE TYPE journal_entry_type AS ENUM ('credit', 'debit');
CREATE TYPE coa_type AS ENUM ('asset', 'liability', 'equity', 'revenue', 'expensis');

CREATE TABLE IF NOT EXISTS users(
    id UUID DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    email_address VARCHAR(255) NOT NULL UNIQUE,
    phone_number VARCHAR(20) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

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

CREATE TABLE IF NOT EXISTS account_blocks(
    id UUID DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    account_id UUID NOT NULL,
    is_current_block BOOLEAN NOT NULL,
    block_size INTEGER NOT NULL,
    transactions_count INTEGER NOT NULL DEFAULT 0,
    status block_status NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS block_meta(
    id UUID DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    account_id UUID NOT NULL,
    block_tx_limit INTEGER NOT NULL,
    transaction_tx_id UUID NOT NULL,
    opening_date TEXT NOT NULL,
    closing_date TEXT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS journal_entries(
    id UUID DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name TEXT NOT NULL,
    type journal_entry_type NOT NULL,
    amount FLOAT NOT NULL,
    block_id UUID NOT NULL,
    transaction_id UUID NOT NULL,
    owner_id UUID NULL,
    memo TEXT NULL,
    account_number TEXT NULL,
    organization_id UUID NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS ledger_accounts(
    id UUID DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    account_number TEXT NOT NULL,
    status ledger_account_status NOT NULL,
    book ledger_book NOT NULL,
    particular TEXT NOT NULL,
    owner_id UUID NULL,
    label TEXT NULL,
    current_active_block_id UUID NOT NULL,
    block_count INTEGER NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS ledger_transactions(
    id UUID DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    description TEXT NOT NULL,
    status transaction_status NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS wallet_type(
    id UUID DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name TEXT NOT NULL,
    rules JSON NOT NULL,
    owner_id UUID NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS wallet(
    id UUID DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name TEXT NOT NULL,
    type UUID NOT NULL,
    account_number VARCHAR(50) NOT NULL,
    ledger_accounts JSON NOT NULL,
    owner_id UUID NOT NULL,
    PRIMARY KEY (id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS chart_of_accounts;
DROP TABLE IF EXISTS account_blocks;
DROP TABLE IF EXISTS block_meta;
DROP TABLE IF EXISTS journal_entries;
DROP TABLE IF EXISTS ledger_accounts;
DROP TABLE IF EXISTS ledger_transactions;
DROP TABLE IF EXISTS wallet;
DROP TABLE IF EXISTS wallet_type;
DROP TYPE IF EXISTS coa_type;
DROP TYPE IF EXISTS block_status;
DROP TYPE IF EXISTS ledger_account_status;
DROP TYPE IF EXISTS transaction_status;
DROP TYPE IF EXISTS ledger_book;
DROP TYPE IF EXISTS journal_entry_type;
-- +goose StatementEnd
