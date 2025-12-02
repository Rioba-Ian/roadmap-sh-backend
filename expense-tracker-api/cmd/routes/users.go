package routes

import (
	"database/sql"
	"net/http"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/controllers"
	"github.com/Rioba-Ian/expense-tracker-api/cmd/middlewares"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterUser() *http.ServeMux {
	r := http.NewServeMux()

	r.Handle("GET /", middlewares.Authenticate(http.HandlerFunc(controllers.GetUsers)))
	r.HandleFunc("POST /signup/", controllers.CreateUser)
	r.HandleFunc("POST /login/", controllers.Login)

	return r
}
