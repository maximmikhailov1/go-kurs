package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	initializers2 "github.com/maximmikhailov1/go-kurs/api/initializers"
	"github.com/maximmikhailov1/go-kurs/api/routes"
	"log"
)

func init() {
	initializers2.LoadEnvVariables()
	initializers2.ConnectToDB()
}

// TODO: Сделать чтобы адрес можно было указать
func main() {
	engine := html.New("./public/assets/views", ".tmpl")
	app := fiber.New(fiber.Config{
		StrictRouting: true,
		Views:         engine,
	})
	routes.SetupRoutes(app)
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
