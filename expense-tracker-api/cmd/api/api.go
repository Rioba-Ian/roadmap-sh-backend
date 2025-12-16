package api

import (
	"database/sql"
	"fmt"
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
	handler := routes.NewHandler(s.DB)

	userHandler := handler.RegisterUser()
	expenseHandler := handler.RegisterExpenses()

	router := http.NewServeMux()
	router.Handle("/users/", http.StripPrefix("/users", userHandler))
	router.Handle("/expenses/", http.StripPrefix("/expenses", expenseHandler))

	return http.ListenAndServe(fmt.Sprintf(":%s", s.addr), router)
}
