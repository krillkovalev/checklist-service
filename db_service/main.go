package main

import (
	"db_service/server"
	"context"
	"log"
	"db_service/config"
	// "os"
)

func main() {
	database, err := config.ConnectPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()


	ctx := context.Background() 
	redis := config.RedisConnection(ctx)

	defer redis.Close()

	// logName := "DBServiceLogs.json"
	// file, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()



	grpcServer := server.NewGRPCServer(":8181")
	grpcServer.Run(redis, ctx, database)
	if err != nil {
		log.Fatalf("problem starting server: %v", err)
	}

	
}