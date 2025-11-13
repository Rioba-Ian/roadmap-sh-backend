package api

import (
	"net/http"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/routes"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{
		addr: addr,
	}
}

func (s *ApiServer) Run() error {
	userHandler := routes.NewHandler()

	userRouter := userHandler.RegisterUser()
	expenseRoutes := routes.RegisterExpenses()

	router := http.NewServeMux()
	router.Handle("/users/", http.StripPrefix("/users", userRouter))
	router.Handle("/expenses/", http.StripPrefix("/expenses", expenseRoutes))

	return http.ListenAndServe(s.addr, router)
}
