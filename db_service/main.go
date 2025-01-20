// TO DO Сделать кеширование данных, добавить .env файл и брать все credentials оттуда, поразбираться с dependency injection

package main

import (
	"context"
	"db_service/config"
	"db_service/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
 
	r := chi.NewRouter()
	r.Use(middleware.Logger)

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
	r.Post("/create", taskHandler.CreateTask)
	r.Put("/done", taskHandler.UpdateTask)
	r.Delete("/delete", taskHandler.DeleteTask)

	return r
}