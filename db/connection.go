package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN string = os.Getenv("DSN")
var DB *gorm.DB

func DBConnection() error {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("DB connected successfully")
	return nil
}
