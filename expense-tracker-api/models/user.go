package models

import "time"

type User struct {
	ID            string    `json:"id"`
	First_name    *string   `json:"first_name"`
	Last_name     *string   `json:"last_name"`
	Password      *string   `json:"password"`
	Email         *string   `json:"email"`
	Token         *string   `json:"token"`
	Refresh_token *string   `json:"refresh_token"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"created_at"`
}
