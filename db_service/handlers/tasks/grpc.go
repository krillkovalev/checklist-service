package handlers

import (
	"context"
	"db_service/generated/tasks"
	"db_service/models"
	"google.golang.org/protobuf/types/known/wrapperspb"

	grpc "google.golang.org/grpc"
)

type TasksGrpcHandler struct {
	dbService models.DBService
	tasks.UnimplementedDBServiceServer
}

func NewGrpcDBService(grpc *grpc.Server, dbService models.DBService) {
	gRPCHandler := &TasksGrpcHandler{
		dbService: dbService,
	}

	tasks.RegisterDBServiceServer(grpc, gRPCHandler)
}

func (h *TasksGrpcHandler) CreateTask(ctx context.Context, req *tasks.CreateTaskRequest) (*wrapperspb.BoolValue, error) {
	task := &models.Task{
		Title: req.Title,
		Body:  req.Body,
	}

	err := h.dbService.CreateTask(ctx, task)
	if err != nil {
		return wrapperspb.Bool(false), err
	}

	return wrapperspb.Bool(true), nil
}
