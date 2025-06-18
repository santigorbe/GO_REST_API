package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/santigorbe/rest_golang/db"
	_ "github.com/santigorbe/rest_golang/docs"
	"github.com/santigorbe/rest_golang/models"
	"github.com/santigorbe/rest_golang/routes"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Mi API en Go
// @version 1.0
// @description Esta es una API REST construida con Go y Gorilla Mux
// @host localhost:3000
// @BasePath /api
// @schemes http
func main() {
	log.SetPrefix("[API-GO] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Println(".env loading error:", err)
	}

	if err := db.DBConnection(); err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Task{})

	router := routes.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Go API!"))
	}).Methods("GET")

	router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/index.html", http.StatusMovedPermanently)
	})

	handler := cors.Default().Handler(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Println("No PORT environment variable detected, using default port 3000")
	} else {
		log.Println("Using PORT environment variable:", port)
	}

	log.Println("Starting server on port " + port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
