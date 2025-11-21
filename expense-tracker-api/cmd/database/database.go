package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connString := os.Getenv("GOOSE_DBSTRING")

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	fmt.Println("Successfully connected to the database!")

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println(rows.Next())
}
