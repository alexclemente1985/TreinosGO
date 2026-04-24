-- +goose Up
CREATE INDEX idx_user_public_id ON users(public_id);
ALTER TABLE users ADD CONSTRAINT public_id_user_length CHECK (char_length(public_id) = 26);


-- +goose Down
ALTER TABLE users DROP CONSTRAINT public_id_user_length;
DROP INDEX idx_user_public_id;
