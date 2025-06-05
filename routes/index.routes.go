package routes

import (
	"github.com/gorilla/mux"
	"github.com/santigorbe/rest_golang/middlewares"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Middlewares
	r.Use(middlewares.RecoverMiddleware)
	r.Use(middlewares.LoggerMiddleware)
	api := r.PathPrefix("/api").Subrouter()

	// Rutas
	RegisterUsersRoutes(api)
	RegisterTasksRoutes(api)

	return r
}
