package config

import (
	"fmt"
	"os"

	"afthab/github.com/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnect() {
	var err error
	dns := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("Error Connecting to the base")
	}
	DB.AutoMigrate(
		models.Task{},
	)
}
