package routes

import (
	"net/http"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/controllers"
	"github.com/Rioba-Ian/expense-tracker-api/cmd/middlewares"
)

func RegisterExpenses() *http.ServeMux {
	r := http.NewServeMux()

	r.Handle("GET /{id}", middlewares.Authenticate(
		http.HandlerFunc(controllers.GetExpense),
	))

	r.Handle("GET /", middlewares.Authenticate(
		http.HandlerFunc(controllers.GetExpenses),
	))

	r.Handle("POST /", middlewares.Authenticate(
		http.HandlerFunc(controllers.CreateExpense),
	))

	r.Handle("DELETE /{id}", middlewares.Authenticate(
		http.HandlerFunc(controllers.DeleteExpense),
	))

	return r
}
