package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
    "api_service/utils"
)

type Task struct {
	Client 		*http.Client
}

type RequestBody struct {
	ID	string `json:"id"`
}

type Messsage struct {
    Timestamp   string  `json:"timestamp"`
    Action      string  `json:"action"`
}

func (t *Task) Create(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("http://localhost:8181/tasks/create", "application/json", r.Body)
	if err != nil {
		log.Fatalf("error in db_service: %v", err) 
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("response status is incorrect: %v", err)
	}

	w.WriteHeader(http.StatusOK)
}	

func (t *Task) List(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get("http://localhost:8181/tasks/list")
	if err != nil {
		log.Fatalf("error in db_service: %v", err)
	}
    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("problem unmarshaling response: %v", err)
    }

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("response status is incorrect: %v", err)
	}

	w.WriteHeader(http.StatusOK)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responseBody)
    
}

func (t *Task) ActiveTasks(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get("http://localhost:8181/tasks/active")
	if err != nil {
		log.Fatalf("error in db_service: %v", err)
	}
    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("problem unmarshaling response: %v", err)
    }

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
	}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responseBody)
}


func (t *Task) DeleteByID(w http.ResponseWriter, r *http.Request) {
	dbReq := RequestBody{}
    if err := json.NewDecoder(r.Body).Decode(&dbReq); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }

	url := fmt.Sprintf("http://localhost:8181/tasks/delete?id=%s", dbReq.ID)
    responseBody, err := utils.ProxyRequest(t.Client, "DELETE", url, dbReq)
    if err != nil {
        log.Fatalf("something wrong with request: %v", err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responseBody)
}



func (t *Task) DoneByID(w http.ResponseWriter, r *http.Request) {
    dbReq := RequestBody{}
    if err := json.NewDecoder(r.Body).Decode(&dbReq); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }
	url := fmt.Sprintf("http://localhost:8181/tasks/done?id=%s", dbReq.ID)
    responseBody, err := utils.ProxyRequest(t.Client, "PUT", url, dbReq)
    if err != nil {
        log.Fatalf("something wrong with request: %v", err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responseBody)
}

