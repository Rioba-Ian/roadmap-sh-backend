package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/service"
	"github.com/Rioba-Ian/expense-tracker-api/models"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(string)

	if !ok {
		http.Error(w, "userId not found in context", http.StatusInternalServerError)
		return
	}

	expenses, err := service.UserExpenses(userID)

	if err != nil {
		http.Error(w, "could not retrieve expenses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println("new expense:: ", newExpense.ExpenseDate.Format("2025-12-02"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(newExpense)

	createdExpense, err := service.CreateExpense(newExpense, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdExpense)
}
