package models

import (
	"database/sql"
	"fmt"
)

func CreateTaskDB(db *sql.DB, title, body string) error {
	_, err := db.Exec("insert into tasks(task_title, task_body) values($1,$2)", title, body)
	if err != nil {
		fmt.Errorf("error inserting data: %v", err)
		return err
	}
	return nil
}

func GetTasksDB(db *sql.DB) ([]Task, error){
	rows, err := db.Query("select * from tasks")
	result := []Task{}
	if err != nil {
		fmt.Errorf("error fetching tasks from db: %v", err)
	} else {
		for rows.Next() {
			tmp := Task{}
			rows.Scan(&tmp.ID, &tmp.Title, &tmp.Body, &tmp.Done)
			result = append(result, tmp)
		}
	}
	return result, nil
}

func DeleteTaskDB(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM tasks WHERE id=$1", id)
 	if err != nil {
		fmt.Errorf("error deleting task from db: %v", err)
	} 
}

func MarkTaskDoneDB(db *sql.DB) {
	_, err := db.Exec("update tasks set done=true")
	if err != nil {
		fmt.Errorf("error updating table in db: %v", err)
	} 
}