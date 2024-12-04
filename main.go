package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/maximmikhailov1/go-kurs/initializers"
	"github.com/maximmikhailov1/go-kurs/routes"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
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
