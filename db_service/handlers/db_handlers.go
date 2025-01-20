package handlers

import (
	"context"
	"database/sql"
	"db_service/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/redis/go-redis/v9"
)

type TaskHandler struct {
	DB 		*sql.DB
	Client	*redis.Client
	Context	context.Context
}

func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {

	task := models.Task{}
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Printf("Unable to create task: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := models.CreateTaskDB(t.DB, task.Title, task.Body)
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

func (t *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request)  {

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

func (t *TaskHandler) ActiveTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	// Используем SCAN для поиска ключей
	iter := t.Client.Scan(t.Context, 0, "task:id:*", 0).Iterator()
	for iter.Next(t.Context) {
		res := t.Client.HGetAll(t.Context, iter.Val())
		if res.Err() != nil {
			http.Error(w, fmt.Sprintf("Redis error: %v", res.Err()), http.StatusInternalServerError)
			return
		}

		tmp := models.Task{}
		if err := res.Scan(&tmp); err != nil {
			http.Error(w, fmt.Sprintf("Failed to parse task from Redis: %v", err), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, tmp)
	}

	// Проверяем ошибки итератора
	if err := iter.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Redis iteration error: %v", err), http.StatusInternalServerError)
		return
	}

	// Если в Redis ничего не найдено, загружаем данные из базы данных
	if len(tasks) == 0 {
		var err error
		tasks, err = models.GetTasksDB(t.DB)
		if err != nil {
			http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
			return
		}
	}

	// Формируем JSON-ответ
	tasksJson, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tasksJson)
}



func (t *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	str := r.URL.Query().Get("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Error to convert query string: %v", err)
	}

	models.DeleteFromCache(t.Context, t.Client, id)
	log.Println(id)
	if err != nil {
		log.Printf("The key has deleted yet: %v\n", err)
	}

	err = models.DeleteTaskDB(t.DB, id)
	if err != nil {
		log.Fatalf("Unable to delete task: %v", err)
	}

	w.WriteHeader(http.StatusOK)
}
func (t *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	str := r.URL.Query().Get("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Error to convert query string: %v", err)
	}

	models.DeleteFromCache(t.Context, t.Client, id)
	if err != nil {
		log.Printf("The key has deleted yet: %v\n", err)
	}

	err = models.MarkTaskDoneDB(t.DB, id)
	if err != nil {
		log.Fatalf("Unable to delete task: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	

}