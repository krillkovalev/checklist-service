package handlers

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	body := []byte(`{
		"title": "Unit test example",
		"body": "here is some test text"
	}`)
	request, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(body))

	response := httptest.NewRecorder()

	task := Task{
		Ctx: context.Background(),
	}

	task.Create(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	actual := response.Body.String()

	contains := `{"id":`
	assert.Contains(t, actual, contains)
}

func TestCreateNegative(t *testing.T) {
	body := []byte(`{
		"title": "Unit test example
		"body": "here is some test text"
	}`)
	request, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(body))

	response := httptest.NewRecorder()

	task := Task{
		Ctx: context.Background(),
	}

	task.Create(response, request)

	if status := response.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	actual := response.Body.String()

	expected := "Bad Request\n"
	assert.Equal(t, expected, actual)
}

func TestList(t *testing.T) {
	request, _ := http.NewRequest("GET", "/getAllTasks", http.NoBody)

	response := httptest.NewRecorder()

	task := Task{
		Ctx: context.Background(),
	}

	task.List(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	actual := response.Body.String()

	assert.Contains(t, actual, "id")
	assert.Contains(t, actual, "body")
	assert.Contains(t, actual, "title")
	assert.Contains(t, actual, "done")
}

func TestActiveTasks(t *testing.T) {
	request, _ := http.NewRequest("GET", "/getActiveTasks", http.NoBody)

	response := httptest.NewRecorder()

	task := Task{
		Ctx: context.Background(),
	}

	task.ActiveTasks(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	actual := response.Body.String()

	assert.Contains(t, actual, "id")
	assert.Contains(t, actual, "body")
	assert.Contains(t, actual, "title")
	assert.Contains(t, actual, "done")
}

func TestUpdateTask(t *testing.T) {

	body := []byte(`{
		"id": 12
	}`)

	request, _ := http.NewRequest("PUT", "/done", bytes.NewBuffer(body))

	response := httptest.NewRecorder()

	task := Task{
		Ctx: context.Background(),
	}

	task.DoneByID(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	actual := response.Body.String()

	assert.Contains(t, actual, "success")
}

func TestUpdateTaskNegative(t *testing.T) {

	body := []byte(`{
		"id": 34356
	}`)

	request, _ := http.NewRequest("PUT", "/done", bytes.NewBuffer(body))

	response := httptest.NewRecorder()

	task := Task{
		Ctx: context.Background(),
	}

	task.DoneByID(response, request)

	if status := response.Code; status != http.StatusBadGateway {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadGateway)
	}

	actual := response.Body.String()

	expected := "Bad Gateway\n"

	assert.Contains(t, actual, expected)
}

func TestUpdateTaskIncorrect(t *testing.T) {

	body := []byte(`{
		"id": fjmkiprdmg21-390o
	}`)

	request, _ := http.NewRequest("PUT", "/done", bytes.NewBuffer(body))

	response := httptest.NewRecorder()

	task := Task{
		Ctx: context.Background(),
	}

	task.DoneByID(response, request)

	if status := response.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	actual := response.Body.String()

	expected := "Bad Request\n"

	assert.Contains(t, actual, expected)
}

func TestDeleteTask(t *testing.T) {

	body := []byte(`{
		"id": 9
	}`)

	request, _ := http.NewRequest("DEL", "/delete", bytes.NewBuffer(body))

	response := httptest.NewRecorder()

	task := Task{
		Ctx: context.Background(),
	}

	task.DeleteByID(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	actual := response.Body.String()

	assert.Contains(t, actual, "success")
}

func TestDeleteTaskNegative(t *testing.T) {

	body := []byte(`{
		"id": 4325430
	}`)

	request, _ := http.NewRequest("DEL", "/delete", bytes.NewBuffer(body))

	response := httptest.NewRecorder()

	task := Task{
		Ctx: context.Background(),
	}

	task.DeleteByID(response, request)

	if status := response.Code; status != http.StatusBadGateway {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadGateway)
	}

	actual := response.Body.String()

	expected := "Bad Gateway\n"
	assert.Equal(t, expected, actual)
}

func TestDeleteTaskIncorrect(t *testing.T) {

	body := []byte(`{
		"id": rekwlmgfre
	}`)

	request, _ := http.NewRequest("DEL", "/delete", bytes.NewBuffer(body))

	response := httptest.NewRecorder()

	task := Task{
		Ctx: context.Background(),
	}

	task.DeleteByID(response, request)

	if status := response.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	actual := response.Body.String()

	expected := "Bad Request\n"
	assert.Equal(t, expected, actual)
}
