package db

import (
	"go-casbin/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(){
	Instance, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})

	if err != nil {
		log.Fatalln("::: :: DB connection error", err)
	}
	log.Println("::: :: Connected to DB")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("::: :: DB Migration Completed")
}