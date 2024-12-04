package main

import (
	"github.com/maximmikhailov1/go-kurs/initializers"
	"github.com/maximmikhailov1/go-kurs/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Driver{},
		&models.Car{},
		&models.Client{},
		&models.Order{})
	if err != nil {
		log.Fatalf("Failed to migrate: %s", err.Error())
	}
}
