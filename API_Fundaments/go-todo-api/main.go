package main

import (
	"encoding/json"
	"net/http"
	"time"
	"os"
)

func loadTasks() {
	file, err := os.Open("tasks.json")
	if err != nil {
		return
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&tasks)
}

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	IsDone      bool   `json:"is_done"`
}

var tasks []Task

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	task.ID = len(tasks) + 1
	task.CreatedAt = time.Now().Format(time.RFC3339)

	tasks = append(tasks, task)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getTasksHandler(w, r)
	} else if r.Method == http.MethodPost {
		createTaskHandler(w, r)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	loadTasks()
	http.HandleFunc("/tasks", tasksHandler)
	http.ListenAndServe(":8080", nil)
}