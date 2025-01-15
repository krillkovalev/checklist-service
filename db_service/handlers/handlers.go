package handlers

import (
	"database/sql"
	"db_service/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	DB *sql.DB
}

func (t TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	task := models.Task{}
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("bad request: %v", err)
		return
	}
	err = models.CreateTaskDB(t.DB, task.Title, task.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request)  {
	tasks, err := models.GetTasksDB(t.DB)
	if err != nil {
		log.Fatalf("Unable to fetch all tasks: %v", err)
	}

	tasksJson, err := json.Marshal(tasks)
	if err != nil {
		log.Fatalf("Unable to form response with list of tasks: %v", err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tasksJson)

}
func (t TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Fatalf("Can't get id parameter: %v", err)
	}

	err = models.DeleteTaskDB(t.DB, id)
	if err != nil {
		log.Fatalf("Unable to delete task: %v", err)
	}

	w.WriteHeader(http.StatusNoContent)
}
func (t TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Fatalf("Can't get id parameter: %v", err)
	}

	err = models.MarkTaskDoneDB(t.DB, id)
	if err != nil {
		log.Fatalf("Unable to delete task: %v", err)
	}

	w.WriteHeader(http.StatusNoContent)
	

}