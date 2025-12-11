package routes

import (
	"database/sql"
	"net/http"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/controllers"
	"github.com/Rioba-Ian/expense-tracker-api/cmd/middlewares"
	"github.com/Rioba-Ian/expense-tracker-api/cmd/service"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		DB: db,
	}
}

func (h *Handler) RegisterUser() *http.ServeMux {
	r := http.NewServeMux()

	userService := service.NewUserService(h.DB)
	userController := controllers.NewUserController(userService)

	r.Handle("GET /", middlewares.Authenticate(
		http.HandlerFunc(userController.GetUsers)),
	)
	r.HandleFunc("POST /signup/",
		userController.CreateUser,
	)
	r.HandleFunc("POST /login/",
		userController.Login,
	)

	return r
}
