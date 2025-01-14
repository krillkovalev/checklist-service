package handlers

import (
	"fmt"
	"net/http"
)

type Task struct {}

func (t *Task) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a task")
}

func (t *Task) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List of tasks")
}

func (t *Task) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete task by ID")
}

func (t *Task) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a task by ID")
}

func (t *Task) DoneByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update task and assign status done for task")
}