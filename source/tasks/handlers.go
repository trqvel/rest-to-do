package tasks

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// GetTasksHandler - обработчик для получения задач
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	tasks, err := GetTasksByUser(userID)
	if err != nil {
		http.Error(w, "Could not fetch tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// CreateTaskHandler - обработчик для создания задачи
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = CreateTask(task.UserID, task.Description)
	if err != nil {
		http.Error(w, "Could not create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
