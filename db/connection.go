package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() error {
	DSN := "postgres://santi:mysecretpassword@db:5432/gorm?sslmode=disable"
	var err error
	log.Println("Connecting to DB: ", DSN)

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("DB connected successfully")
	return nil
}
