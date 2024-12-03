package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maximmikhailov1/go-kurs/initializers"
	"github.com/maximmikhailov1/go-kurs/models"
	"github.com/maximmikhailov1/go-kurs/utils"
	"math/rand/v2"
	"net/http"
)

func OrderTaxiCreateAndRender(c *fiber.Ctx) error {
	//fmt.Printf("Cookies: %v", c.Cookies("auth"))
	var clientData fiber.Map
	clientData = utils.ParseClientJWT(c.Cookies("authClient"))
	var order models.Order
	order.ClientID = uint(clientData["ID"].(float64))

	var driverID uint
	var drivers []models.Driver
	result := initializers.DB.Find(&drivers)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to find any drivers",
			"error":   result.Error.Error(),
		})
	}
	driverID = rand.UintN(uint(len(drivers)))
	assignedDriver := drivers[driverID]
	order.DriverID = assignedDriver.ID

	var car models.Car
	err := initializers.DB.Model(&assignedDriver).Association("Car").Find(&car)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to find a car of a driver",
			"error":   err,
		})
	}
	result = initializers.DB.Create(&order)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to create an order",
			"error":   result.Error.Error(),
		})
	}
	return c.Render("orderTaxi", fiber.Map{
		"DriverName":  assignedDriver.FirstName,
		"CarFirmName": car.FirmName,
		"CarModel":    car.ModelName,
		"CarRegPlate": car.RegPlateNumber,
		"CarColor":    car.Color,
	})
}
