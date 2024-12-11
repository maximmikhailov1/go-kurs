package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maximmikhailov1/go-kurs/api/initializers"
	models2 "github.com/maximmikhailov1/go-kurs/api/models"
	"net/http"
	"strconv"
)

func OrderRender(c *fiber.Ctx) error {
	return c.Render("orders", fiber.Map{})
}
func OrderIndex(c *fiber.Ctx) error {
	var orders []models2.Order
	result := initializers.DB.Find(&orders)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to find any orders",
			"error":   result.Error.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "found orders successfully",
		"orders":  orders,
	})
}
func OrderShow(c *fiber.Ctx) error {
	id := c.Params("id")

	var order models2.Order
	result := initializers.DB.First(&order, id)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to find a order with given id",
			"error":   result.Error.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("found order id:%v successfully", id),
		"order":   order,
	})
}
func OrderDriverShow(c *fiber.Ctx) error {
	driverIdString := c.Params("driverId")
	driverId, err := strconv.ParseUint(driverIdString, 10, 0)
	var clients []models2.Client
	err = initializers.DB.Joins("Orders", initializers.DB.Where(&models2.Order{DriverID: uint(driverId)})).Select("first_name").Find(&clients).Error
	//err = initializers.DB.Find(&clients, "id IN ?", orders).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to find clients of orders of the driver",
			"error":   err.Error(),
		})
	}
	log.Info(clients)
	return c.Status(http.StatusOK).JSON(clients)
}

//func OrderCreate(c *fiber.Ctx) error {
//	var order models.Order
//	err := c.BodyParser(&order)
//	if err != nil {
//		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
//			"message": "failed to parse form to order",
//			"error":   err.Error(),
//		})
//	}
//	result := initializers.DB.Create(&order)
//	if result.Error != nil {
//		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
//			"message": "failed to create an order",
//			"error":   result.Error.Error(),
//		})
//	}
//	return c.Status(http.StatusOK).JSON(fiber.Map{
//		"message": "order created",
//	})
//}
