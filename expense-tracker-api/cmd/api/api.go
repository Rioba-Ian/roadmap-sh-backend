package api

import (
	"database/sql"
	"net/http"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/routes"
)

type ApiServer struct {
	addr string
	DB   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		DB:   db,
	}
}

func (s *ApiServer) Run() error {
	userHandler := routes.NewHandler(s.DB)
	expenseHandler := routes.NewHandler(s.DB)

	userRouter := userHandler.RegisterUser()
	expenseRoutes := expenseHandler.RegisterExpenses()

	router := http.NewServeMux()
	router.Handle("/users/", http.StripPrefix("/users", userRouter))
	router.Handle("/expenses/", http.StripPrefix("/expenses", expenseRoutes))

	return http.ListenAndServe(s.addr, router)
}
