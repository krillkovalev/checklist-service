package handlers

import (
	"api_service/models"
	"api_service/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Task struct {
	Client 		*http.Client
}

func (t *Task) Create(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Post("http://localhost:8181/tasks/create", "application/json", r.Body)
	if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
        return
	}

    record := models.Messsage{
        Timestamp: time.Now().Format("2006-01-02 15:04:05"),
        Action: "create",
    }

    msg, err := json.Marshal(record)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    err = models.PushMessageToQueue("tasks-log-topic", msg)
    if err != nil {
        http.Error(w, "Bad Gateway", http.StatusBadGateway)
        return
    }

	w.WriteHeader(http.StatusOK)
}	

func (t *Task) List(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get("http://localhost:8181/tasks/list")
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
        return
	}
    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
        return
	}

    record := models.Messsage{
        Timestamp: time.Now().Format("2006-01-02 15:04:05"),
        Action: "list",
    }

    msg, err := json.Marshal(record)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    err = models.PushMessageToQueue("tasks-log-topic", msg)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responseBody)
    
}

func (t *Task) ActiveTasks(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get("http://localhost:8181/tasks/active")
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
        return
	}
    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
	}

    record := models.Messsage{
        Timestamp: time.Now().Format("2006-01-02 15:04:05"),
        Action: "active",
    }

    msg, err := json.Marshal(record)
    if err != nil {
        http.Error(w, "Bad Gateway", http.StatusBadGateway)
        return
    }

    err = models.PushMessageToQueue("tasks-log-topic", msg)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responseBody)
}


func (t *Task) DeleteByID(w http.ResponseWriter, r *http.Request) {
	dbReq := models.RequestBody{}
    if err := json.NewDecoder(r.Body).Decode(&dbReq); err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

	url := fmt.Sprintf("http://localhost:8181/tasks/delete?id=%s", dbReq.ID)
    responseBody, err := utils.ProxyRequest(t.Client, "DELETE", url, dbReq)
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    record := models.Messsage{
        Timestamp: time.Now().Format("2006-01-02 15:04:05"),
        Action: "deletion",
    }

    msg, err := json.Marshal(record)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    err = models.PushMessageToQueue("tasks-log-topic", msg)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responseBody)
}



func (t *Task) DoneByID(w http.ResponseWriter, r *http.Request) {
    dbReq := models.RequestBody{}
    if err := json.NewDecoder(r.Body).Decode(&dbReq); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	url := fmt.Sprintf("http://localhost:8181/tasks/done?id=%s", dbReq.ID)
    responseBody, err := utils.ProxyRequest(t.Client, "PUT", url, dbReq)
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    record := models.Messsage{
        Timestamp: time.Now().Format("2006-01-02 15:04:05"),
        Action: "done",
    }

    msg, err := json.Marshal(record)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    err = models.PushMessageToQueue("tasks-log-topic", msg)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responseBody)
}

