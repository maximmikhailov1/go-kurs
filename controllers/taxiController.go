package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maximmikhailov1/go-kurs/initializers"
	"github.com/maximmikhailov1/go-kurs/models"
	"github.com/maximmikhailov1/go-kurs/utils"
	"gorm.io/gorm"
	"math/rand/v2"
	"net/http"
)

func OrderTaxiCreateAndRender(c *fiber.Ctx) error {
	//fmt.Printf("Cookies: %v", c.Cookies("auth"))
	var clientData fiber.Map
	clientData = utils.ParseJWT(c.Cookies("authClient"))
	var order models.Order
	order.ClientID = uint(clientData["ID"].(float64))

	var driverID uint
	var drivers []models.Driver
	result := initializers.DB.Not("car_id IS NULL").Find(&drivers)
	query := initializers.DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Not("drivers.car_id = ?", nil).Find(&drivers)
	})
	log.Info(query)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to find any drivers",
			"error":   result.Error.Error(),
		})
	}
	log.Info(drivers, '\n')
	driverID = uint(rand.IntN(len(drivers)))
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
	log.Info(assignedDriver)
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
