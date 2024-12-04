package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maximmikhailov1/go-kurs/controllers"
	"github.com/maximmikhailov1/go-kurs/middleware"
	"github.com/maximmikhailov1/go-kurs/utils"
)

func SetupRoutes(app *fiber.App) {
	app.Static("/static", "./public/assets")
	app.Use(middleware.Authorize)
	app.Use(middleware.IsAuthenticated)
	app.Get("/", func(c *fiber.Ctx) error {
		loggedIn := false
		if _, ok := c.Locals("loggedIn").(bool); ok {
			loggedIn = true
		}
		return c.Render("index", fiber.Map{"LoggedIn": loggedIn})
	})
	app.Get("/orders", controllers.OrderRender)
	app.Get("/order-taxi", controllers.OrderTaxiCreateAndRender)
	app.Get("/auth/client", controllers.AuthClientRender)

	app.Get("/auth/driver", controllers.AuthDriverRender)

	app.Get("/drivers/car", controllers.DriverCarSelectionRender)

	app.Post("/api/driver/sign-in", controllers.SignInDriver)
	app.Post("/api/driver/sign-up", controllers.SignUpDriver)

	app.Post("/api/client/sign-in", controllers.SignInClient)
	app.Post("/api/client/sign-up", controllers.SignUpClient)
	app.Get("/api/logout", controllers.Logout)

	app.Get("/api/orders", controllers.OrderIndex)
	app.Get("/api/orders/:id", controllers.OrderShow)

	app.Post("/api/cars", controllers.CarCreate)
	app.Get("/api/cars", controllers.CarsIndex)
	app.Get("/api/cars/not-used", controllers.CarsNotUsedIndex)
	app.Get("/api/cars/:id", controllers.CarShow)
	app.Put("/api/cars/:id", controllers.CarUpdate)
	app.Delete("/api/cars/:id", controllers.CarDelete)

	app.Post("/api/drivers", controllers.DriverCreate)
	app.Get("/api/drivers", controllers.DriversIndex)
	app.Get("/api/drivers/get-id", utils.DriverGetId)
	app.Get("/api/drivers/:id/car", controllers.DriverCarShow)
	app.Put("/api/drivers/:id/car", controllers.DriverCarUpdate)
	app.Get("/api/drivers/:id", controllers.DriverShow)

	app.Put("/api/drivers/:id", controllers.DriverUpdate)

	app.Delete("/api/drivers/:id", controllers.DriverDelete)

	app.Get("/api/orders/driver/:driverId", controllers.OrderDriverShow)

}
