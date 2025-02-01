package service

import (
	"context"
	"database/sql"
	"db_service/models"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"
	// "time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	db *sql.DB
	mock sqlmock.Sqlmock
	err error
)

func TestMain(m *testing.M) {
	fmt.Println("testing start")
	fmt.Println("============================")

	var dbMock *sql.DB
	dbMock, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db = dbMock
	defer db.Close()

	code := m.Run()

	fmt.Println("============================")
	fmt.Println("testing end")
	os.Exit(code)
}

func TestCreateTaskDB(t *testing.T) {
	query := regexp.QuoteMeta("insert into tasks(task_title, task_body) values(?, ?) returning id")
	mock.ExpectExec(query).WithArgs("Db test mock", "Mock test's body")

	task := &models.Task{
		Title: "Db test mock",
		Body: "Mock test's body",
	}

	id, err := CreateTaskDB(context.Background(), task, db)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet(), id)
}

