package service

import (
	"log"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/database"
	"github.com/Rioba-Ian/expense-tracker-api/models"
)

func GetUser(userId string) (*models.User, error) {
	db := database.GetDB()
	var user models.User

	// query := "SELECT first_name, last_name, email FROM users WHERE id = ?"
	row := db.QueryRow("SELECT id, first_name, email FROM users WHERE id = $1", userId)
	if err := row.Scan(&user.ID, &user.First_name, &user.Email); err != nil {
		log.Printf("error scanning row for getting user: ", err.Error())
		return nil, err
	}

	return &user, nil
}

func GetUserExpenses() error {
	return nil
}

func UpdateUserTokens(token, refreshToken, userId string) error {
	db := database.GetDB()

	row := db.QueryRow("UPDATE users SET token = $1, refresh_token = $2 WHERE id = $3", token, refreshToken, userId)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

// func UpdateUser(userId string) (*models.User, error) {
// 	db := database.GetDB()
// 	var user models.User

// }
