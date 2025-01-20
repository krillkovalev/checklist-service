package models

import (
	"database/sql"
	"fmt"
)

func CreateTaskDB(db *sql.DB, title, body string) (int, error) {
	var id int
	query := "insert into tasks(task_title, task_body) values($1, $2) returning id"	
	err := db.QueryRow(query, title, body).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("error creating task in db: %v", err)
	}
	return id, nil
}

func GetTasksDB(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("select * from tasks")
	result := []Task{}
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks from db: %v", err)
	} else {
		for rows.Next() {
			tmp := Task{}
			rows.Scan(&tmp.ID, &tmp.Title, &tmp.Body, &tmp.Done)
			result = append(result, tmp)
		}
	}
	return result, nil
}

func GetActiveTasksDB(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("select * from tasks where done = 'false'")
	result := []Task{}
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks from db: %v", err)
	} else {
		for rows.Next() {
			tmp := Task{}
			rows.Scan(&tmp.ID, &tmp.Title, &tmp.Body, &tmp.Done)
			result = append(result, tmp)
		}
	}
	return result, nil
}

func DeleteTaskDB(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id=$1", id)
 	if err != nil {
		return fmt.Errorf("error deleting task from db: %v", err)
	}
	return nil 
}

func MarkTaskDoneDB(db *sql.DB, id int) error{
	_, err := db.Exec("update tasks set done=true where id=$1", id)
	if err != nil {
		return fmt.Errorf("error updating table in db: %v", err)
	} 
	return nil
}