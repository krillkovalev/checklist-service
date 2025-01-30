package config

import (
	pb "api_service/generated/tasks"
	"log"
	"os"

	"github.com/joho/godotenv"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitClient() (pb.DBServiceClient, *grpc.ClientConn, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Failed to load .env file")
	}

	address := os.Getenv("DB_SERVICE_URL")
	if address == "" {
		address = "localhost:8181"
	}

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	return pb.NewDBServiceClient(conn), conn, nil
}
