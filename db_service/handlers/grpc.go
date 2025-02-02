package handlers

import (
	"context"
	"database/sql"
	pb "db_service/generated/tasks"
	"db_service/models"
	"db_service/service"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TasksGrpcHandler struct {
	Client *redis.Client
	DBConn *sql.DB
	pb.UnimplementedDBServiceServer
}

func NewGrpcDBService(grpc *grpc.Server, client *redis.Client, ctx context.Context, db *sql.DB) {
	gRPCHandler := &TasksGrpcHandler{
		Client: client,
		DBConn: db,
	}
	pb.RegisterDBServiceServer(grpc, gRPCHandler)
}

func (t *TasksGrpcHandler) CreateTask(ctx context.Context, in *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	task := &models.Task{
		Title: in.Title,
		Body:  in.Body,
	}

	id, err := service.CreateTaskDB(task, t.DBConn)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error retrieving tasks from the db")
	}
	task.ID = int64(id)
	task.Done = false
	key := fmt.Sprintf(service.KeyFormat, task.ID)
	if err := service.ToRedisSet(ctx, t.Client, key, task); err != nil {
		log.Printf("Error caching in redis: %v", err)
	}

	return &pb.CreateTaskResponse{Id: int64(task.ID)}, nil
}

func (t *TasksGrpcHandler) ListTasks(ctx context.Context, in *pb.EmptyRequest) (*pb.ListTasksResponse, error) {
	
	tasks, err := service.GetTasksDB(t.DBConn)
	if err != nil {
		log.Printf("Unable to fetch all tasks: %v", err)
		return nil, status.Errorf(codes.NotFound, "error retrieving tasks from the db")
	}

	response := pb.ListTasksResponse{
		Tasks: make([]*pb.Task, len(tasks)),
	}

	for index, task := range tasks {
		response.Tasks[index] = &pb.Task{
			Id: int64(task.ID),
			Title: task.Title,
			Body:  task.Body,
			Done:  task.Done,
		}
	}

	return &response, nil

}

func (t *TasksGrpcHandler) ActiveTasks(ctx context.Context, in *pb.EmptyRequest) (*pb.ListTasksResponse, error) {
	errorChan := make(chan error)

	go func () {
		_, err := service.UpdateCache(t.DBConn, t.Client)
		errorChan <- err
		close(errorChan)
	}()

	var tasks []models.Task
	var err error
	
	err = <- errorChan
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Failed to fetch caching data")
	}

	iter := t.Client.Scan(ctx, 0, "task:id:*", 0).Iterator()
	for iter.Next(ctx) {
		res := t.Client.HGetAll(ctx, iter.Val())
		if res.Err() != nil {
			return nil, status.Errorf(codes.Internal, "Failed to fetch caching data")
		}

		tmp := models.Task{}
		if err := res.Scan(&tmp); err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to parse task from redis")
		}
		tasks = append(tasks, tmp)
	}

	// Проверяем ошибки итератора
	if err := iter.Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "Redis iteration error")
	}

	response := pb.ListTasksResponse{
		Tasks: make([]*pb.Task, len(tasks)),
	}

	for index, task := range tasks {
		response.Tasks[index] = &pb.Task{
			Id: int64(task.ID),
			Title: task.Title,
			Body:  task.Body,
			Done:  task.Done,
		}
	}

	return &response, nil
}

func (t *TasksGrpcHandler) DeleteTask(ctx context.Context, in *pb.DeleteTaskRequest) (*pb.SuccessResponse, error) {
	id := in.Id

	err := service.DeleteFromCache(ctx, t.Client, id)
	if err != nil {
		log.Printf("The key has deleted yet: %v\n", err)
	}

	err = service.DeleteTaskDB(t.DBConn, int(id))
	if err != nil {

		return &pb.SuccessResponse{Body: "fail"}, status.Errorf(codes.Internal, "Unable to delete task")
	}
	return &pb.SuccessResponse{Body: "success"}, nil
}


func (t *TasksGrpcHandler) UpdateTask(ctx context.Context, in *pb.UpdateTaskRequest) (*pb.SuccessResponse, error) {
	id := in.Id

	err := service.DeleteFromCache(ctx, t.Client, id)
	if err != nil {
		log.Printf("The key has deleted yet: %v\n", err)
	}

	err = service.MarkTaskDoneDB(t.DBConn, int(id))
	if err != nil {
		return &pb.SuccessResponse{Body: "fail"}, status.Errorf(codes.Internal, "Unable to update task")
	}
	return &pb.SuccessResponse{Body: "success"}, nil
}