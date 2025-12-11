-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD blacklist TEXT;

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_expenses_id ON expenses(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN blacklist;
-- +goose StatementEnd
