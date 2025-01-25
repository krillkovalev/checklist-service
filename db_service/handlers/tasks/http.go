package handlers

import (
	"context"
	"db_service/models"
	"db_service/service"
	"log"
	"encoding/json"
	"net/http"
	"services/common/genproto/db_service"
	"github.com/go-chi/chi/v5"
)

type TaskHttpHandler struct {
	dbService models.DBService
}

func NewHttpDBService(dbService models.DBService) *TaskHttpHandler{
	handler := &TaskHttpHandler{
		dbService: dbService,
	}

	return handler
}

func (t *TaskHttpHandler) CreateTask(w http.ResponseWriter, r *http.Request) {

	task := models.Task{}
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Printf("Unable to create task: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := service.CreateTaskDB(t.DB, task.Title, task.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error querying in db: %v", err)
		return
	}

	task.ID = id
	task.Done = false
	key := fmt.Sprintf(models.KeyFormat, task.ID)
	if err := models.ToRedisSet(t.Context, t.Client, key, &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error caching in redis: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TasksHttpHandler) RegisterRouter(router *chi.Mux) {
	router.Post("/create", h.dbService.CreateTask)
}