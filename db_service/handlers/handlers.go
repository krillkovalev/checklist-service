// To do: Дописать хендлеры

package handlers

import (
	"database/sql"
	"db_service/models"
	"encoding/json"
	"log"
	"net/http"
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
	}
	err = models.CreateTaskDB(t.DB, task.Title, task.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}



}
func (t TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request)  {}
func (t TaskHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {}
func (t TaskHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {}