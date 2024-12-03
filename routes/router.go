package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maximmikhailov1/go-kurs/controllers"
	"github.com/maximmikhailov1/go-kurs/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Static("/static", "./public/assets")
	app.Use(middleware.Authorize)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
	app.Get("/orders", controllers.OrderRender)
	app.Get("/order-taxi", controllers.OrderTaxiCreateAndRender)
	app.Get("/authClient", controllers.AuthRender)
	app.Post("/api/client/sign-in", controllers.SignInClient)
	app.Post("/api/client/sign-up", controllers.SignUpClient)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/orders", controllers.OrderIndex)
	app.Get("/api/orders/:id", controllers.OrderShow)

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
