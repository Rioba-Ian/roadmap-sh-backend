package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/service"
	"github.com/Rioba-Ian/expense-tracker-api/helpers"
	"github.com/Rioba-Ian/expense-tracker-api/models"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(u *service.UserService) *UserController {
	return &UserController{
		UserService: u,
	}
}

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(string)

	if !ok {
		http.Error(w, "userId not found in context", http.StatusInternalServerError)
		return
	}

	fmt.Printf("userid:: %s\n", userID)
	log.Println("userId:: %s", userID)
	user, err := c.UserService.GetUser(userID)
	if err != nil {
		http.Error(w, "could not find user details", http.StatusInternalServerError)
		return
	}

	userRes := user.ToUserPublic()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userRes)
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if newUser.FirstName == "" || newUser.Password == "" || newUser.Email == "" {
		http.Error(w, "first name, password and email are required", http.StatusBadRequest)
		return
	}

	if err := helpers.CheckPasswordStrength(newUser.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword := helpers.HashPassword(&newUser.Password)
	newUser.PasswordHash = *hashedPassword

	db := c.UserService.DB

	var err error
	dbQuery := `INSERT INTO users (first_name, email, password_hash) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`

	row := db.QueryRow(dbQuery, newUser.FirstName, newUser.Email, newUser.PasswordHash)
	if err = row.Scan(&newUser.ID, &newUser.CreatedAt, &newUser.UpdatedAt); err != nil {
		log.Printf("error scanning row: %v", err)
	}
	if err != nil {
		log.Printf("Database insert error:: %v", err)
		http.Error(w, "Failed to create user", http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := helpers.GenerateTokens(newUser.Email, newUser.ID)

	if err != nil {
		log.Printf("failed to create user tokens:: %v", err)
		http.Error(w, "Failed to create user tokens", http.StatusInternalServerError)
		return
	}

	newUser.Token = accessToken
	newUser.RefreshToken = refreshToken
	userRes := newUser.ToUserPublic()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userRes)
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var foundUser models.User
	db := c.UserService.DB

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == "" {
		http.Error(w, "email and password is required", http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT id, email, password_hash FROM users WHERE email = $1", user.Email)
	if err := row.Scan(&foundUser.ID, &foundUser.Email, &foundUser.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, "Unknown error occured", http.StatusInternalServerError)
			return
		}
	}

	passwordIsValid, err := helpers.VerifyPassword(foundUser.PasswordHash, user.Password)
	if !passwordIsValid {
		http.Error(w, "invalid email/password", http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Printf("error occured in verifying password %w", err)
		http.Error(w, "unexpected error occurred", http.StatusInternalServerError)
		return
	}

	accessToken, refreshToken, err := helpers.GenerateTokens(foundUser.Email, foundUser.ID)

	if err != nil {
		log.Printf("failed to create user tokens:: %v", err)
		http.Error(w, "Failed to create user tokens", http.StatusInternalServerError)
		return
	}

	foundUser.Token = accessToken
	foundUser.RefreshToken = refreshToken
	err = c.UserService.UpdateUserTokens(accessToken, refreshToken, foundUser.ID)

	if err != nil {
		log.Printf("failed to update user tokens:: %v", err)
		http.Error(w, "Failed to update user tokens", http.StatusInternalServerError)
		return
	}

	userRes := foundUser.ToUserPublic()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// TODO: convert found user object to public user object
	json.NewEncoder(w).Encode(userRes)
}

func (c *UserController) UpdateUserDetails(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// var foundUser models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
