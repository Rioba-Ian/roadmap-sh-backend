package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/database"
	"github.com/Rioba-Ian/expense-tracker-api/helpers"
	"github.com/Rioba-Ian/expense-tracker-api/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get users")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if newUser.First_name == "" || newUser.Password == "" || newUser.Email == "" {
		http.Error(w, "first name, password and email are required", http.StatusBadRequest)
		return
	}

	hashedPassword := helpers.HashPassword(&newUser.Password)
	newUser.PasswordHash = *hashedPassword

	db := database.GetDB()

	var err error
	dbQuery := `INSERT INTO users (first_name, email, password_hash) VALUES ($1, $2, $3) RETURNING id, inserted_at, updated_at`

	row := db.QueryRow(dbQuery, newUser.First_name, newUser.Email, newUser.PasswordHash)
	if err = row.Scan(&newUser.ID, &newUser.Created_at, &newUser.Updated_at); err != nil {
		log.Printf("error scanning row", err.Error())
	}
	if err != nil {
		log.Printf("Database insert error:: %w", err)
		http.Error(w, "Failed to create user", http.StatusBadRequest)
		return
	}

	accessToken, refreshToken := helpers.GenerateTokens(newUser.Email, newUser.ID)

	newUser.Token = accessToken
	newUser.Refresh_token = refreshToken
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
