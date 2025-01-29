package service

import (
	"context"
	"database/sql"
	"db_service/models"
	"fmt"
	"errors"
)

func CreateTaskDB(ctx context.Context, task *models.Task, db *sql.DB) (int, error) {
	var id int
	query := "insert into tasks(task_title, task_body) values($1, $2) returning id"	
	err := db.QueryRow(query, task.Title, task.Body).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("error creating task in db: %v", err)
	}
	return id, nil
}

func GetTasksDB(db *sql.DB) ([]models.Task, error) {
	rows, err := db.Query("select * from tasks")
	result := []models.Task{}
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks from db: %v", err)
	} else {
		for rows.Next() {
			tmp := models.Task{}
			rows.Scan(&tmp.ID, &tmp.Title, &tmp.Body, &tmp.Done)
			result = append(result, tmp)
		}
	}
	return result, nil
}

func GetActiveTasksDB(db *sql.DB) ([]models.Task, error) {
	rows, err := db.Query("select * from tasks where done = 'false'")
	result := []models.Task{}
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks from db: %v", err)
	} else {
		for rows.Next() {
			tmp := models.Task{}
			rows.Scan(&tmp.ID, &tmp.Title, &tmp.Body, &tmp.Done)
			result = append(result, tmp)
		}
	}
	return result, nil
}

func DeleteTaskDB(db *sql.DB, id int) error {
	res, err := db.Exec("DELETE FROM tasks WHERE id=$1", id)
 	if err != nil {
		return fmt.Errorf("error deleting task from db: %v", err)
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("error deleting nonexisting item")
	}
	return nil
	
}

func MarkTaskDoneDB(db *sql.DB, id int) error {
	res, err := db.Exec("update tasks set done=true where id=$1", id)
	if err != nil {
		return fmt.Errorf("error updating table in db: %v", err)
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("error updating nonexisting item")
	} 
	return nil
}