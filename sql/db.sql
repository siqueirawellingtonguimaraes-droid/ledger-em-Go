CREATE TABLE IF NOT EXISTS accounts (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
);

CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(255) PRIMARY KEY,
    description TEXT,
    credit_account_id VARCHAR(255) NOT NULL,
    debit_account_id VARCHAR(255) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    date TIMESTAMP NOT NULL,
    FOREIGN KEY (credit_account_id) REFERENCES accounts(id),
    FOREIGN KEY (debit_account_id) REFERENCES accounts(id)
);