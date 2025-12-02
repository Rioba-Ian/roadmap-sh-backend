package service

import "github.com/Rioba-Ian/expense-tracker-api/cmd/database"

func GetUser() error {
	return nil
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
