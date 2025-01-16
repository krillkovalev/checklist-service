package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Task struct {
	Client 		*http.Client
}

type RequestBody struct {
	ID	string `json:"id"`
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



func (t *Task) DeleteByID(w http.ResponseWriter, r *http.Request) {
	dbReq := RequestBody{}
    if err := json.NewDecoder(r.Body).Decode(&dbReq); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }

	url := fmt.Sprintf("http://localhost:8181/tasks/delete?id=%s", dbReq.ID)
    responseBody, err := t.proxyRequest("DELETE", url, dbReq)
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
    responseBody, err := t.proxyRequest("PUT", url, dbReq)
    if err != nil {
        log.Fatalf("something wrong with request: %v", err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responseBody)
}

func (t *Task) proxyRequest(method, url string, body interface{}) ([]byte, error) {
    // Преобразуем тело запроса в JSON
    jsonBytes, err := json.Marshal(body)
    if err != nil {
        return nil, fmt.Errorf("error marshaling request body: %v", err)
    }

    // Создаем HTTP-запрос
    req, err := http.NewRequest(method, url, bytes.NewReader(jsonBytes))
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %v", err)
    }

    // Отправляем запрос
    resp, err := t.Client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to send request: %v", err)
    }
    defer resp.Body.Close()

    // Проверяем статус-код ответа
    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
    }

    // Читаем тело ответа
    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %v", err)
    }

    return responseBody, nil
}