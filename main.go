package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/santigorbe/rest_golang/db"
	_ "github.com/santigorbe/rest_golang/docs" // ← esta es tu carpeta generada
	"github.com/santigorbe/rest_golang/models"
	"github.com/santigorbe/rest_golang/routes"
	httpSwagger "github.com/swaggo/http-swagger"
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

	db.DB.AutoMigrate(models.Task{}, models.User{})

	router := routes.NewRouter()

	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)

	handler := cors.Default().Handler(router)
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, handler)
	log.Println("Visit http://localhost:" + port)
	log.Println("Visit http://localhost:" + port + "/docs for API documentation")
}
