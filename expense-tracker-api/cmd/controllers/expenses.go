package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(string)

	if !ok {
		http.Error(w, "userId not found in context", http.StatusInternalServerError)
		return
	}

	log.Println("userId:: %s", userID)

	fmt.Fprintf(w, "get expenses")
}
