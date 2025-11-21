package controllers

import (
	"fmt"
	"net/http"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get expenses")
}
