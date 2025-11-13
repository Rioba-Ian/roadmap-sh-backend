package routes

import "net/http"

func RegisterExpenses() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Still in progress"))
	})

	return r
}
