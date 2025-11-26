package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connString := os.Getenv("GOOSE_DBSTRING")

	DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(0)

	if err = DB.Ping(); err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	fmt.Println("Successfully connected to the database!")

	// rows, err := db.Query("SELECT * FROM users")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// fmt.Println(rows.Next())
}

func GetDB() *sql.DB {
	if DB == nil {
		log.Fatal("Database not initialized.")
	}
	return DB
}

func CloseDB() {
	if DB == nil {
		return
	}

	if err := DB.Close(); err != nil {
		log.Fatalf("error closing db conn: %w", err)
	}

	log.Println("Database connection closed.")

}
