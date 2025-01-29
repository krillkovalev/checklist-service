package server

import (
	"db_service/handlers"
	"github.com/redis/go-redis/v9"
	"database/sql"
	"context"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
} 

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr:  addr}
}

func (s *gRPCServer) Run(client *redis.Client, ctx context.Context, db *sql.DB) error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	handlers.NewGrpcDBService(grpcServer, client, ctx, db)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}