package service

import (
	"context"
	"database/sql"
	"fmt"
	"services/common/genproto/db_service"
	"github.com/redis/go-redis/v9"
)

type TaskHandler struct {
	DB 		*sql.DB
	Client	*redis.Client
	Context	context.Context
}

func NewDBService() *TaskHandler {
	return &TaskHandler{}
}

func (t *TaskHandler) CreateTaskDB(ctx context.Context, task *db_service.Task) (int, error) {
	var id int
	query := "insert into tasks(task_title, task_body) values($1, $2) returning id"	
	err := t.DB.QueryRow(query, task.Title, task.Body).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("error creating task in db: %v", err)
	}
	return id, nil
}