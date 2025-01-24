package main

import (
	"api_service/config"
	"api_service/handlers"
	"api_service/middlewares"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
)

func main() {

	logName := "ApiServiceLogs.json"
	file, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := config.CreateLogger(logName)

	r := chi.NewRouter()
	r.Use(middlewares.Logger(&logger))

	taskClient := handlers.Task{Client: &http.Client{}}
	r.Mount("/api_service", TaskRoutes(taskClient))

	fmt.Println("Server api_service is running")
	http.ListenAndServe(":8282", r)
}

func TaskRoutes(taskClient handlers.Task) chi.Router {
	r := chi.NewRouter()

	r.Get("/getAllTasks", taskClient.List)
	r.Get("/getActiveTasks", taskClient.ActiveTasks)
	r.Post("/create", taskClient.Create)
	r.Put("/done", taskClient.DoneByID)
	r.Delete("/delete", taskClient.DeleteByID)

	return r
}
