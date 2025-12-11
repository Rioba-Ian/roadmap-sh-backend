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
	userService := service.NewUserService(h.DB)
	expenseController := controllers.NewExpenseController(expenseService)
	middlewareHandler := middlewares.NewMiddleWare(userService)

	r.Handle("GET /{id}", middlewareHandler.Authenticate(
		http.HandlerFunc(expenseController.GetExpense),
	))

	r.Handle("GET /", middlewareHandler.Authenticate(
		http.HandlerFunc(expenseController.GetExpenses),
	))

	r.Handle("POST /", middlewareHandler.Authenticate(
		http.HandlerFunc(expenseController.CreateExpense),
	))

	r.Handle("DELETE /{id}", middlewareHandler.Authenticate(
		http.HandlerFunc(expenseController.DeleteExpense),
	))

	return r
}
