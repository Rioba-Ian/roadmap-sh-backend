-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD blacklist TEXT;

CREATE INDEX idx_users_email ON users(email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN blacklist;
DROP INDEX IF EXISTS idx_users_email;
-- +goose StatementEnd
