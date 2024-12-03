package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/maximmikhailov1/go-kurs/initializers"
	"github.com/maximmikhailov1/go-kurs/models"
	"net/http"
)

func OrderRender(c *fiber.Ctx) error {
	return c.Render("order", fiber.Map{})
}
func OrderIndex(c *fiber.Ctx) error {
	var orders []models.Order
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

	var order models.Order
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
