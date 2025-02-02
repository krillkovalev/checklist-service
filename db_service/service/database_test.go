package service

import (
	"database/sql"
	"db_service/models"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"errors"
)

var (
	db *sql.DB
	mock sqlmock.Sqlmock
	err error
)

func generateTask() models.Task {
	return models.Task {
	 Title:       fmt.Sprintf("title %d", rand.Intn(9999)),
	 Body: fmt.Sprintf("description %d", rand.Intn(9999)),
	}
   }

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

	task := &models.Task{
		Title: "Db test mock",
		Body:  "Mock test's body",
	}

	query := regexp.QuoteMeta("insert into tasks(task_title, task_body) values($1, $2) returning id")

	mock.ExpectQuery(query).
		WithArgs(task.Title, task.Body).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(42))

	id, err := CreateTaskDB(task, db)

	assert.NoError(t, err)
	assert.NotZero(t, id) 
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateTaskDBNegative(t *testing.T) {

	task := &models.Task{
		Title: "Db negatuve test mock",
		Body:  "Mock negative test's body",
	}

	query := regexp.QuoteMeta("select * from tasks")

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("error creating task in db"))
	
	id, err := CreateTaskDB(task, db)

	assert.Error(t, err)
	assert.Zero(t, id)
	assert.Contains(t, err.Error(), "error creating task in db") 
}

func TestGetTasksDB(t *testing.T) {

	query := regexp.QuoteMeta("select * from tasks")
	var tasks []models.Task
	task := generateTask()
	tasks = append(tasks, task)

	testCases := []struct {
		name          string
		expected      []models.Task
		expectedQuery string
		expectedErr   error
	   }{
		{
		 name:          "success",
		 expected:      tasks,
		 expectedQuery: query,
		 expectedErr:   nil,
		},
	   }

	for _, test := range testCases {

		Expectedrows := sqlmock.NewRows([]string{"id", "task_title", "task_body", "done"})

		for _, task := range test.expected {
			Expectedrows.AddRow(task.ID, task.Title, task.Body, task.Done)
		}

		mock.ExpectQuery(test.expectedQuery).WillReturnRows(Expectedrows)
		
		tasks, err := GetTasksDB(db)

		assert.Len(t, tasks, len(test.expected))
		assert.Equal(t, test.expected, tasks)
		assert.Equal(t, test.expectedErr, err) 
	}
}

func TestGetTasksDBNegative(t *testing.T) {

	query := regexp.QuoteMeta("select * from tasks")

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("error fetching tasks from db:"))
	
	tasks, err := GetTasksDB(db)

	assert.Error(t, err)
	assert.Nil(t, tasks)
	assert.Contains(t, err.Error(), "error fetching tasks from db:") 
}

func TestGetActiveTasksDB(t *testing.T) {

	query := regexp.QuoteMeta("select * from tasks where done = 'false'")
	var tasks []models.Task
	task := generateTask()
	tasks = append(tasks, task)

	testCases := []struct {
		name          string
		expected      []models.Task
		expectedQuery string
		expectedErr   error
	   }{
		{
		 name:          "success",
		 expected:      tasks,
		 expectedQuery: query,
		 expectedErr:   nil,
		},
	   }

	for _, test := range testCases {

		Expectedrows := sqlmock.NewRows([]string{"id", "task_title", "task_body", "done"})

		for _, task := range test.expected {
			Expectedrows.AddRow(task.ID, task.Title, task.Body, task.Done)
		}

		mock.ExpectQuery(test.expectedQuery).WillReturnRows(Expectedrows)
		
		tasks, err := GetActiveTasksDB(db)

		assert.Len(t, tasks, len(test.expected))
		assert.Equal(t, test.expected, tasks)
		assert.Equal(t, test.expectedErr, err) 
	}
}

func TestGetActiveTasksDBNegative(t *testing.T) {

	query := regexp.QuoteMeta("select * from tasks where done = 'false'")

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("error fetching tasks from db:"))
	
	tasks, err := GetActiveTasksDB(db)

	assert.Error(t, err)
	assert.Nil(t, tasks)
	assert.Contains(t, err.Error(), "error fetching tasks from db:") 
}

func TestDeleteTaskDB(t *testing.T) {

	query := regexp.QuoteMeta("DELETE FROM tasks WHERE id=$1")

	mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	err := DeleteTaskDB(db, 1)

	assert.NoError(t, err) 
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteDBNegative(t *testing.T) {

	query := regexp.QuoteMeta("DELETE FROM tasks WHERE id=$1")

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("error deleting task from db:"))
	
	err := DeleteTaskDB(db, 1)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error deleting task from db:") 
}

func TestMarkTaskDoneDB(t *testing.T) {

	query := regexp.QuoteMeta("update tasks set done=true where id=$1")

	mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	err := MarkTaskDoneDB(db, 1)

	assert.NoError(t, err) 
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMarkTaskDoneDBNegative(t *testing.T) {

	query := regexp.QuoteMeta("update tasks set done=true where id=$1")

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("error updating table in db:"))
	
	err := MarkTaskDoneDB(db, 1)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error updating table in db:") 
}
