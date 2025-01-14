package config

import (
	"database/sql"
	"fmt"
)

func connectPostgresDB() *sql.DB {
	connstr := "user=krillkovalev dbname=checklist_info password='108814' host=localhost port=5432 sslmode-disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		fmt.Println(err)
	}
	return db
}	