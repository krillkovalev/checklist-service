package routes

import (
	"net/http"
	"api-service/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/tasks", loadTaskRoutes)

	return router
}

func loadTaskRoutes(router chi.Router) {
	taskHandler := &handlers.Task{}

	router.Post("/", taskHandler.Create)
	router.Get("/", taskHandler.List)
	router.Get("/", taskHandler.GetByID)
	router.Put("/{id}", taskHandler.DoneByID)
	router.Delete("/{id}", taskHandler.DeleteByID)
}