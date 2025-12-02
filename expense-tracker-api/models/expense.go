package models

import "time"

type Expense struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Amount      string    `json:"amount"`
	Descripiton string    `json:"description"`
	ExpenseDate time.Time `json:"expense_date"`
}
