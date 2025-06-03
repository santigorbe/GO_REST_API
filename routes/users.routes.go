package routes

import (
	"github.com/gorilla/mux"
	"github.com/santigorbe/rest_golang/controllers"
	"github.com/santigorbe/rest_golang/repository"
	"github.com/santigorbe/rest_golang/services"
)

func RegisterUsersRoutes(r *mux.Router) {
	userRepo := &repository.GormUserRepository{}
	userService := &services.UserService{Repo: userRepo}
	userHandler := &controllers.UserHandler{Service: userService}

	r.HandleFunc("/users", userHandler.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", userHandler.GetUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", userHandler.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/user/{id}", userHandler.UpdateUserHandler).Methods("PUT")
}
