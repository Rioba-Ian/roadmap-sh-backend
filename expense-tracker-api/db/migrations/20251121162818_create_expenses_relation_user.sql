-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION set_updated_at_to_now()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = (now() at time zone 'utc');
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

create table expenses (
  id uuid primary key default gen_random_uuid(),
  user_id uuid not null,
  amount decimal(10, 2) not null,
  description text,
  expense_date date not null,
  created_at timestamp with time zone  default (now() at time zone 'utc'),
  updated_at timestamp with time zone  default (now() at time zone 'utc')
);
-- Foreign key
ALTER TABLE expenses
ADD CONSTRAINT fk_user
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

-- create trigger
CREATE TRIGGER set_updated_at_expenses
BEFORE UPDATE ON expenses
FOR EACH ROW
EXECUTE FUNCTION set_updated_at_to_now();

CREATE TRIGGER set_updated_at_users
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION set_updated_at_to_now();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE expenses DROP CONSTRAINT fk_user;
DROP TABLE IF EXISTS expenses;
DROP FUNCTION IF EXISTS set_updated_at_to_now() CASCADE;
-- +goose StatementEnd
