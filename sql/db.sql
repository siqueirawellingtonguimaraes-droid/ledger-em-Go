CREATE TABLE IF NOT EXISTS accounts (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    currency TEXT
);

CREATE TABLE IF NOT EXISTS transactions (
    id TEXT PRIMARY KEY,
    description TEXT,
    credit_account_id TEXT NOT NULL,
    debit_account_id TEXT NOT NULL,
    amount INTEGER NOT NULL,
    date TIMESTAMP NOT NULL,
    FOREIGN KEY (credit_account_id) REFERENCES accounts(id),
    FOREIGN KEY (debit_account_id) REFERENCES accounts(id)
);