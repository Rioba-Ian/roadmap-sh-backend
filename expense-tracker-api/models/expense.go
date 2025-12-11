package models

import "time"

type Expense struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Amount      string    `json:"amount"`
	Description string    `json:"description"`
	ExpenseDate time.Time `json:"expense_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
