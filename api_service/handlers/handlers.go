package handlers

import (
	pb "api_service/generated/tasks"
	"api_service/utils"
	"context"
	"net/http"
)

type Task struct {
	Ctx context.Context
}

func (t *Task) Create(w http.ResponseWriter, r *http.Request) {
	utils.HandleRequest(w, r, "create", &pb.CreateTaskRequest{}, func(client pb.DBServiceClient, ctx context.Context, req any) (any, error) {
		return client.CreateTask(ctx, req.(*pb.CreateTaskRequest))
	})
}

func (t *Task) List(w http.ResponseWriter, r *http.Request) {
	utils.HandleRequest(w, r, "list", &pb.EmptyRequest{}, func(client pb.DBServiceClient, ctx context.Context, req any) (any, error) {
		return client.ListTasks(ctx, req.(*pb.EmptyRequest))
	})
}

func (t *Task) ActiveTasks(w http.ResponseWriter, r *http.Request) {
	utils.HandleRequest(w, r, "active", &pb.EmptyRequest{}, func(client pb.DBServiceClient, ctx context.Context, req any) (any, error) {
		return client.ActiveTasks(ctx, req.(*pb.EmptyRequest))
	})
}

func (t *Task) DeleteByID(w http.ResponseWriter, r *http.Request) {
	utils.HandleRequest(w, r, "deletion", &pb.DeleteTaskRequest{}, func(client pb.DBServiceClient, ctx context.Context, req any) (any, error) {
		return client.DeleteTask(ctx, req.(*pb.DeleteTaskRequest))
	})
}

func (t *Task) DoneByID(w http.ResponseWriter, r *http.Request) {
	utils.HandleRequest(w, r, "done", &pb.UpdateTaskRequest{}, func(client pb.DBServiceClient, ctx context.Context, req any) (any, error) {
		return client.UpdateTask(ctx, req.(*pb.UpdateTaskRequest))
	})
}
