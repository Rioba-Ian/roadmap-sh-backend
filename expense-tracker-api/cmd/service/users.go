package service

import (
	"database/sql"
	"log"

	"github.com/Rioba-Ian/expense-tracker-api/models"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (s *UserService) GetUser(userId string) (*models.User, error) {
	db := s.DB
	var user models.User

	// query := "SELECT first_name, last_name, email FROM users WHERE id = ?"
	row := db.QueryRow("SELECT id, first_name, email FROM users WHERE id = $1", userId)
	if err := row.Scan(&user.ID, &user.FirstName, &user.Email); err != nil {
		log.Printf("error scanning row for getting user: %v", err)
		return nil, err
	}

	return &user, nil
}

func GetUserExpenses() error {
	return nil
}

func (s *UserService) UpdateUserTokens(token, refreshToken, userId string) error {
	db := s.DB

	_, err := db.Exec("UPDATE users SET token = $1, refresh_token = $2 WHERE id = $3", token, refreshToken, userId)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

// func UpdateUser(userId string) (*models.User, error) {
// 	db := database.GetDB()
// 	var user models.User

// }
