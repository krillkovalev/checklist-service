// TO DO Дописать запуск сервера

package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"db_service/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		
	})

	r.Mount("/tasks", TaskRoutes())

	http.ListenAndServe(":8181", r)
}

func TaskRoutes() chi.Router {
	r := chi.NewRouter()
	taskHandler := handlers.TaskHandler{} 

	r.Get("/", taskHandler.ListTasks)
	r.Post("/", taskHandler.CreateBook)
	r.Put("/{id}", taskHandler.UpdateBook)
	r.Delete("/{id}", taskHandler.DeleteBook)

	return r
}