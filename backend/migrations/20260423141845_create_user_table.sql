-- +goose Up
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    public_id CHAR(26) UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS users;
