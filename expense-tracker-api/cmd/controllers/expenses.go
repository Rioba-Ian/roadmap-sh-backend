package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/service"
	"github.com/Rioba-Ian/expense-tracker-api/models"
)

type ExpenseController struct {
	ExpenseService *service.ExpenseService
}

func NewExpenseController(e *service.ExpenseService) *ExpenseController {
	return &ExpenseController{
		ExpenseService: e,
	}
}

func (c *ExpenseController) GetExpenses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(string)

	if !ok {
		http.Error(w, "userId not found in context", http.StatusInternalServerError)
		return
	}

	expenses, err := c.ExpenseService.UserExpenses(userID)

	if err != nil {
		http.Error(w, "could not retrieve expenses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
}

func (c *ExpenseController) GetExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(string)

	if !ok {
		http.Error(w, "userId not found in context", http.StatusInternalServerError)
		return
	}

	expenseId := r.PathValue("id")
	if expenseId == "" {
		http.Error(w, "id not found", http.StatusBadRequest)
		return
	}

	fmt.Println("expense with id::", expenseId)

	expense, err := c.ExpenseService.UserExpenseId(userID, expenseId)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "could not find expense "+expenseId, http.StatusNotFound)
			return
		} else {
			http.Error(w, "error fetching expense ", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&expense)
}

func (c *ExpenseController) CreateExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(string)
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "could not read json data", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if !ok {
		http.Error(w, "userId not found in context", http.StatusInternalServerError)
		return
	}

	var newExpense *models.Expense
	// var err error
	if err := json.Unmarshal(body, &newExpense); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdExpense, err := c.ExpenseService.CreateExpense(newExpense, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdExpense)
}

func (c *ExpenseController) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, ok := ctx.Value("userID").(string)
	expenseId := r.PathValue("id")

	if !ok {
		http.Error(w, "userId not found in context", http.StatusInternalServerError)
		return
	}

	if expenseId == "" {
		http.Error(w, "id not found", http.StatusBadRequest)
		return
	}

	err := c.ExpenseService.DeleteExpense(userId, expenseId)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, "unexpected error occured", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
