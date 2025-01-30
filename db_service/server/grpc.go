package server

import (
	"context"
	"database/sql"
	"db_service/handlers"
	"db_service/middleware"
	"log"
	"net"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/redis/go-redis/v9"

	grpc "google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
} 

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr:  addr}
}

func (s *gRPCServer) Run(client *redis.Client, ctx context.Context, db *sql.DB, file *os.File) error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logger := log.New(file, "", log.LstdFlags|log.Lshortfile)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(middleware.InterceptorLogger(logger)),
		),
	)

	handlers.NewGrpcDBService(grpcServer, client, ctx, db)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}