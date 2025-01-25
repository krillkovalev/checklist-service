package handlers

import (
	"db_service/types"
	"services/common/genproto/db_service"
	"context"
	grpc "google.golang.org/grpc"
)

type TasksGrpcHandler struct {
	dbService types.DBService
	db_service.UnimplementedDBServiceServer
}

func NewGrpcDBService(grpc *grpc.Server, dbService types.DBService) {
	gRPCHandler := &TasksGrpcHandler{
		dbService: dbService,
	}

	db_service.RegisterDBServiceServer(grpc, gRPCHandler)
}

func (h *TasksGrpcHandler) CreateTask(ctx context.Context, req *db_service.CreateTaskRequest) (*db_service.CreateTaskResponse, error) {
	task := &types.Task{
		Title: req.Title,
		Body: req.Body,

	}

	err := h.dbService.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}

	res := &db_service.CreateTaskResponse {
		Status: "success",
	}

	return res, nil
}