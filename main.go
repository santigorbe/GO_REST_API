package main

import (
	"log"
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/santigorbe/rest_golang/db"
	_ "github.com/santigorbe/rest_golang/docs" // ← esta es tu carpeta generada
	"github.com/santigorbe/rest_golang/models"
	"github.com/santigorbe/rest_golang/routes"
	// httpSwagger "github.com/swaggo/http-swagger"
)

// @title Mi API en Go
// @version 1.0
// @description Esta es una API REST construida con Go y Gorilla Mux
// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	log.SetPrefix("[API-GO] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Println(".env loading error:", err)
	}

	db.DBConnection()

	log.Println("Migrating User model")
	db.DB.AutoMigrate(&models.User{})
	log.Println("Migrating Task model")
	db.DB.AutoMigrate(&models.Task{})


	router := routes.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Go API!"))
	}).Methods("GET")

	router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	handler := cors.Default().Handler(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port if not set in .env
		log.Println("No PORT environment variable detected, using default port 3000")
	} else {
		log.Println("Using PORT environment variable:", port)
	}

	log.Println("Starting server on port " + port)
	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	log.Println("Visit http://localhost:" + port)
	log.Println("Visit http://localhost:" + port + "/docs for API documentation")
}
