package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maximmikhailov1/go-kurs/initializers"
	"github.com/maximmikhailov1/go-kurs/routes"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// TODO: Человек пришёл со своей машиной
func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
