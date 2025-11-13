package controllers

import (
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "will display endpoint for all")
}
