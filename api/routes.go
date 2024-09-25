package api

import (
	"github.com/gorilla/mux"
	"github.com/trqvel/rest-to-do/source/tasks"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", tasks.GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks", tasks.CreateTaskHandler).Methods("POST")

	return router
}
