-- +goose Up
CREATE TABLE customers (
    id BIGSERIAL PRIMARY KEY,
    public_id CHAR(26) UNIQUE NOT NULL,
    name TEXT NOT NULL,
    email TEXT,
    phone TEXT,
    mobile_phone TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE accounts (
    id BIGSERIAL PRIMARY KEY,
    public_id CHAR(26) UNIQUE NOT NULL,
    customer_id INT REFERENCES customers(id) ON DELETE CASCADE,
    type TEXT,
    balance DOUBLE PRECISION
);

-- +goose Down
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS customers;
