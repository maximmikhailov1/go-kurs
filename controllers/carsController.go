package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/maximmikhailov1/go-kurs/initializers"
	"github.com/maximmikhailov1/go-kurs/models"
	"net/http"
)

func CarCreate(c *fiber.Ctx) error {
	// Получить машину
	var car models.Car
	err := c.BodyParser(&car)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse form to car model",
			"error":   err.Error(),
		})
	}
	// Создать машину
	result := initializers.DB.Create(&car)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to create a car",
			"error":   result.Error.Error(),
		})
	}
	// Вернуть её
	return c.Status(200).JSON(fiber.Map{
		"message": "car successfully created",
		"car":     car,
	})
}

func CarsNotUsedIndex(c *fiber.Ctx) error {
	var cars []models.Car
	initializers.DB.Find(&cars).Where("is_being_used = ?", false)
	return c.Status(http.StatusOK).JSON(cars)
}

func CarsIndex(c *fiber.Ctx) error {
	var cars []models.Car
	result := initializers.DB.Find(&cars)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to find any cars",
			"error":   result.Error.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "found cars successfully",
		"cars":    cars,
	})
}
func CarShow(c *fiber.Ctx) error {
	//Получить URL с id
	id := c.Params("id")

	//Получить машину с нужным id
	var car models.Car
	result := initializers.DB.First(&car, id)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to find a car with given id",
			"error":   result.Error.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("found car id:%v succesfully", id),
		"car":     car,
	})
}
func CarUpdate(c *fiber.Ctx) error {
	//Получить URL с id
	id := c.Params("id")

	// Получить машину
	var carNew models.Car
	var carOld models.Car
	err := c.BodyParser(&carNew)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse form to car",
			"error":   err.Error(),
		})
	}
	//Получить машину с нужным id
	result := initializers.DB.First(&carOld, id)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("failed to find car with id %v", id),
			"error":   result.Error.Error(),
		})
	}
	result = initializers.DB.Model(&carOld).Updates(&carNew)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("failed to update car with id %v", id),
			"error":   result.Error.Error(),
		})
	}
	// Обновить данные
	//Ответить обновленными данными
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("updated car with id %v successfully", id),
		"car":     carNew,
	})
}
func CarDelete(c *fiber.Ctx) error {
	//Получить id машины
	id := c.Params("id")
	//Удалить машину
	result := initializers.DB.Delete(&models.Car{}, id)
	//Ответить о результате
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to delete car",
			"error":   result.Error.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "car deleted successfully",
	})

}
