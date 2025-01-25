package main

import (
	"db_service/handlers"
	"db_service/service"
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

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// register grpc services

	dbService := service.NewDBService()
	handlers.NewGrpcDBService(grpcServer, dbService)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}