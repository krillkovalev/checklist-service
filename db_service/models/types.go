package models

import (
	"context"
	"services/common/genproto/db_service"
)

type Task struct {
	ID			int			`redis:"id" json:"id"`
	Title		string		`redis:"title" json:"title"`
	Body		string		`redis:"body" json:"body"`
	Done		bool		`redis:"done" json:"done"`
}

type DBService interface {
	CreateTask(ctx context.Context, task *db_service.Task) error
}