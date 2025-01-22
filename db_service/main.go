package main

import (
	"context"
	"db_service/config"
	"db_service/handlers"
	"db_service/middlewares"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
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
	
	logName := "DBServiceLogs.json"
	file, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := config.CreateLogger(logName)

	r := chi.NewRouter()
	r.Use(middlewares.Logger(logger))

	taskHandler := handlers.TaskHandler{DB: database, Client: redis, Context: ctx} 
	r.Mount("/tasks", TaskRoutes(taskHandler))


	fmt.Println("Server db_service is running")
	err = http.ListenAndServe(":8181", r)
	if err != nil {
		log.Fatalf("problem starting server: %v", err)
	}
}

func TaskRoutes(taskHandler handlers.TaskHandler) chi.Router {
	r := chi.NewRouter()

	r.Get("/list", taskHandler.ListTasks)
	r.Get("/active", taskHandler.ActiveTasks)
	r.Post("/create", taskHandler.CreateTask)
	r.Put("/done", taskHandler.UpdateTask)
	r.Delete("/delete", taskHandler.DeleteTask)

	return r
}