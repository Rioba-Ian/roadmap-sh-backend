package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/service"
	"github.com/Rioba-Ian/expense-tracker-api/helpers"
)

type Middleware struct {
	UserService *service.UserService
}

func NewMiddleWare(u *service.UserService) *Middleware {
	return &Middleware{
		UserService: u,
	}
}

func (m *Middleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		claims, err := helpers.ValidateToken(tokenString)
		if err != nil {
			log.Printf("Token validation error: %w\n", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims.UserID)

		// check if token is blacklisted
		blackListed := m.UserService.CheckBlackListed(tokenString, claims.ID)
		if blackListed {
			log.Printf("Token expired: %w\n", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
