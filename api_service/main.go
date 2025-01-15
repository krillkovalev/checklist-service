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
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {	
	})

 
	r.Mount("/api_service", TaskRoutes())


	fmt.Println("Server api_service is running")
	http.ListenAndServe(":8282", r)
}

func TaskRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/list", handlers.ListTasks)
	r.Post("/create", handlers.CreateTask)
	r.Put("/done", handlers.UpdateTask)
	r.Delete("/delete", handlers.DeleteTask)

	return r


}
