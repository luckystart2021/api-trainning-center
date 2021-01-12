package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

const (
	HOST = "localhost"
	PORT = 5432
)

// Initialize executes create connection with postgres db
func Initialize() (*sql.DB, error) {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	username, password, database :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")

	if len(username) == 0 || len(password) == 0 || len(database) == 0 {
		log.Fatalf("Error loading data config")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)
	// Open the connection
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return conn, err
	}

	err = conn.Ping()
	if err != nil {
		return conn, err
	}
	log.Println("Database connection established")
	return conn, nil
}
