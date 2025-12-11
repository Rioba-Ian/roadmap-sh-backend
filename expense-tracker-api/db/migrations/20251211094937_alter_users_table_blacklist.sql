-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD blacklist TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN blacklist;
-- +goose StatementEnd
