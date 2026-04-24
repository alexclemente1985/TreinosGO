-- +goose Up
CREATE INDEX idx_customers_public_id ON customers(public_id);
ALTER TABLE customers ADD CONSTRAINT public_id_length CHECK (char_length(public_id) = 26);

CREATE INDEX idx_account_public_id ON accounts(public_id);
ALTER TABLE accounts ADD CONSTRAINT public_id_account_length CHECK (char_length(public_id) = 26);

-- +goose Down
ALTER TABLE accounts DROP CONSTRAINT public_id_account_length;
DROP INDEX idx_account_public_id;

ALTER TABLE customers DROP CONSTRAINT public_id_length;
DROP INDEX idx_customers_public_id;
