package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectPostgresDB() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
	pgUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	connstr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", pgUser, dbName, dbPass, dbHost, dbPort)
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to db: %v", err)
	}
	return db, nil 
}	