package routes

import (
	"net/http"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/controllers"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterUser() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /", controllers.GetUsers)

	return r
}
