package routes

import (
	"net/http"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/controllers"
	"github.com/Rioba-Ian/expense-tracker-api/cmd/middlewares"
	"github.com/Rioba-Ian/expense-tracker-api/cmd/service"
)

func (h *Handler) RegisterExpenses() *http.ServeMux {
	r := http.NewServeMux()

	expenseService := service.NewExpenseService(h.DB)
	expenseController := controllers.NewExpenseController(expenseService)

	r.Handle("GET /{id}", middlewares.Authenticate(
		http.HandlerFunc(expenseController.GetExpense),
	))

	r.Handle("GET /", middlewares.Authenticate(
		http.HandlerFunc(expenseController.GetExpenses),
	))

	r.Handle("POST /", middlewares.Authenticate(
		http.HandlerFunc(expenseController.CreateExpense),
	))

	r.Handle("DELETE /{id}", middlewares.Authenticate(
		http.HandlerFunc(expenseController.DeleteExpense),
	))

	return r
}
