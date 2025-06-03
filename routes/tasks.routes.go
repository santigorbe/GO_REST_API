package routes

import (
	"github.com/gorilla/mux"
	"github.com/santigorbe/rest_golang/controllers"
	"github.com/santigorbe/rest_golang/repository"
	"github.com/santigorbe/rest_golang/services"
)

func RegisterTasksRoutes(r *mux.Router) {
	taskRepo := &repository.GormTaskRepository{}
	taskService := &services.TaskService{Repo: taskRepo}
	taskHandler := &controllers.TaskHandler{Service: taskService}

	r.HandleFunc("/tasks", taskHandler.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/task/{id}", taskHandler.GetTaskHandler).Methods("GET")
	r.HandleFunc("/task/{id}", taskHandler.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/task/{id}", taskHandler.UpdateTaskHandler).Methods("PUT")
}
