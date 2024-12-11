package main

import (
	initializers2 "github.com/maximmikhailov1/go-kurs/api/initializers"
	models2 "github.com/maximmikhailov1/go-kurs/api/models"
	"log"
)

func init() {
	initializers2.LoadEnvVariables()
	initializers2.ConnectToDB()
}

func main() {
	err := initializers2.DB.AutoMigrate(&models2.Driver{},
		&models2.Car{},
		&models2.Client{},
		&models2.Order{})
	if err != nil {
		log.Fatalf("Failed to migrate: %s", err.Error())
	}
}
