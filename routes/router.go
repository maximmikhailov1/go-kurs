package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maximmikhailov1/go-kurs/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Static("/static", "./public/assets")

	app.Post("/api/cars", controllers.CarCreate)
	app.Get("/api/cars", controllers.CarsIndex)
	app.Get("/api/cars/:id", controllers.CarShow)
	app.Put("/api/cars/:id", controllers.CarUpdate)
	app.Delete("/api/cars/:id", controllers.CarDelete)
	app.Post("/api/drivers", controllers.DriverCreate)
	app.Get("/api/drivers", controllers.DriversIndex)
	app.Get("/api/drivers/:id", controllers.DriverShow)
	app.Put("/api/drivers/:id", controllers.DriverUpdate)
	app.Delete("/api/drivers/:id", controllers.DriverDelete)
}
