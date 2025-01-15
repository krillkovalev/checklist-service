package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

func ConnectPostgresDB() (*sql.DB, error) {
	connstr := "user=krillkovalev dbname=checklist_info password=108814 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to db: %v", err)
	}
	return db, nil 
}	