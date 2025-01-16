package main

import (

	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"api_service/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	taskClient := handlers.Task{Client: &http.Client{}}
	r.Mount("/api_service", TaskRoutes(taskClient))

	fmt.Println("Server api_service is running")
	http.ListenAndServe(":8282", r)
}

func TaskRoutes(taskClient handlers.Task) chi.Router {
	r := chi.NewRouter()

	r.Get("/list", taskClient.List)
	r.Post("/create", taskClient.Create)
	r.Put("/done", taskClient.DoneByID)
	r.Delete("/delete", taskClient.DeleteByID)

	return r
}
