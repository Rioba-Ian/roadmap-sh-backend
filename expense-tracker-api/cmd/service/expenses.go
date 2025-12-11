package service

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/Rioba-Ian/expense-tracker-api/models"
)

type ExpenseService struct {
	DB *sql.DB
}

func NewExpenseService(db *sql.DB) *ExpenseService {
	return &ExpenseService{
		DB: db,
	}
}

func (s *ExpenseService) UserExpenses(userId string) ([]models.Expense, error) {
	db := s.DB
	var expenses []models.Expense

	query := `
select e.id, e.user_id, e.amount, e.description, e.expense_date, e.created_at, e.updated_at
from expenses as e
left join users on e.user_id = users.id
where users.id = $1
	`
	rows, err := db.Query(query, userId)
	if err != nil {
		log.Printf("error getting expenses for user ", err, userId)
		return nil, err
	}
	for rows.Next() {
		var e models.Expense
		// var u models.User
		if err := rows.Scan(&e.ID, &e.UserID, &e.Amount,
			&e.Description, &e.ExpenseDate, &e.CreatedAt, &e.UpdatedAt); err != nil {
			log.Printf("could not return expenses for user %w", err)
			return nil, err
		}
		expenses = append(expenses, e)
	}

	if err := rows.Err(); err != nil {
		return expenses, err
	}

	return expenses, nil
}

func (s *ExpenseService) UserExpenseId(userId, expenseId string) (*models.Expense, error) {
	db := s.DB
	var e models.Expense

	query := `
select e.id, e.user_id, e.amount, e.description, e.expense_date, e.created_at, e.updated_at
from expenses as e
left join users on e.user_id = users.id
where users.id = $1 and e.id = $2
	`
	row := db.QueryRow(query, userId, expenseId)
	if err := row.Err(); err != nil {
		log.Printf("error getting expense for user ", err, userId)
		return nil, err
	}

	if err := row.Scan(&e.ID, &e.UserID,
		&e.Amount, &e.Description, &e.ExpenseDate, &e.CreatedAt, &e.UpdatedAt); err != nil {
		log.Printf("error scanning row for expense details: %v", err)
		return nil, err
	}

	return &e, nil
}

func (s *ExpenseService) CreateExpense(newExpense *models.Expense, userId string) (*models.Expense, error) {
	db := s.DB
	var expense models.Expense

	log.Println("new expense userId", newExpense.ExpenseDate, userId)

	query := `
INSERT INTO expenses (amount, description, expense_date, user_id)
VALUES ($1, $2, $3, $4) RETURNING id, amount, description, expense_date, user_id, created_at, updated_at
	`
	intAmount, err := strconv.ParseFloat(newExpense.Amount, 64)
	if err != nil {
		log.Printf("cannot convert amount to number %v", err)
		return nil, err
	}
	row := db.QueryRow(query, intAmount, newExpense.Description,
		newExpense.ExpenseDate, userId,
	)

	if err := row.Scan(&expense.ID, &expense.Amount, &expense.Description,
		&expense.ExpenseDate, &expense.UserID, &expense.CreatedAt, &expense.UpdatedAt); err != nil {
		log.Printf("error creating new expense", err)
		return nil, err
	}

	return &expense, nil
}

func (s *ExpenseService) DeleteExpense(userId, id string) error {
	db := s.DB
	query := `
	DELETE FROM expenses
	WHERE id = $1 AND user_id = $2
	`

	result, err := db.Exec(query, id, userId)
	if err != nil {
		log.Printf("error deleting expense record", id)
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
