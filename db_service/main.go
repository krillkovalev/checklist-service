// TO DO Дописать запуск сервера

package main

import (
	"db_service/config"
	"db_service/handlers"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	database, err := config.ConnectPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		
	})

	taskHandler := handlers.TaskHandler{DB: database} 
	r.Mount("/tasks", TaskRoutes(taskHandler))


	fmt.Println("Server db_service is running")
	http.ListenAndServe(":8181", r)
}

func TaskRoutes(taskHandler handlers.TaskHandler) chi.Router {
	r := chi.NewRouter()

	r.Get("/list", taskHandler.ListTasks)
	r.Post("/create", taskHandler.CreateTask)
	r.Put("/update/{id}", taskHandler.UpdateTask)
	r.Delete("/delete/{id}", taskHandler.DeleteTask)

	return r
}