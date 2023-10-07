-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE block_status AS ENUM ('open', 'close');
CREATE TYPE ledger_account_status AS ENUM ('pending', 'approved');
CREATE TYPE transaction_status AS ENUM ('pending', 'approved');
CREATE TYPE ledger_book AS ENUM ('general_journal', 'cash_receipt');
CREATE TYPE journal_entry_type AS ENUM ('credit', 'debit');

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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account_blocks;
DROP TABLE IF EXISTS block_meta;
DROP TABLE IF EXISTS journal_entries;
DROP TABLE IF EXISTS ledger_accounts;
DROP TABLE IF EXISTS ledger_transactions;
DROP TYPE IF EXISTS block_status;
DROP TYPE IF EXISTS ledger_account_status;
DROP TYPE IF EXISTS transaction_status;
DROP TYPE IF EXISTS ledger_book;
DROP TYPE IF EXISTS journal_entry_type;
-- +goose StatementEnd
