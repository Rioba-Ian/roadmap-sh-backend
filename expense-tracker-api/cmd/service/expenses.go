package service

import (
	"log"
	"strconv"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/database"
	"github.com/Rioba-Ian/expense-tracker-api/models"
)

func UserExpenses(userId string) ([]models.Expense, error) {
	db := database.GetDB()
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
			&e.Descripiton, &e.ExpenseDate, &e.CreatedAt, &e.UpdatedAt); err != nil {
			log.Fatalln("could not return expenses for user")
			return nil, err
		}
		expenses = append(expenses, e)
	}

	if err := rows.Err(); err != nil {
		return expenses, err
	}

	return expenses, nil
}

func UserExpenseId(userId, expenseId string) (*models.Expense, error) {
	db := database.GetDB()
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
		&e.Amount, &e.Descripiton, &e.ExpenseDate, &e.CreatedAt, &e.UpdatedAt); err != nil {
		log.Printf("error scanning row for expense details", err.Error())
		return nil, err
	}

	return &e, nil
}

func CreateExpense(newExpense *models.Expense, userId string) (*models.Expense, error) {
	db := database.GetDB()
	var expense models.Expense

	log.Println("new expense userId", newExpense.ExpenseDate, userId)

	query := `
INSERT INTO expenses (amount, description, expense_date, user_id)
VALUES ($1, $2, $3, $4) RETURNING id, amount, description, expense_date, user_id, created_at, updated_at
	`
	intAmount, err := strconv.Atoi(newExpense.Amount)
	if err != nil {
		log.Printf("cannot convert amount to number")
	}
	row := db.QueryRow(query, intAmount, newExpense.Descripiton,
		newExpense.ExpenseDate, userId,
	)

	if err := row.Scan(&expense.ID, &expense.Amount, &expense.Descripiton,
		&expense.ExpenseDate, &expense.UserID, &expense.CreatedAt, &expense.UpdatedAt); err != nil {
		log.Printf("error creating new expense", err)
		return nil, err
	}

	return &expense, nil
}

func DeleteExpense(id string) error {
	db := database.GetDB()
	query := `
	DELETE FROM expenses
	WHERE id = $1
	`

	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("error deleting expense record", id)
		return err
	}

	return nil
}
