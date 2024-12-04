package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"

	"github.com/maximmikhailov1/go-kurs/initializers"
	"github.com/maximmikhailov1/go-kurs/models"
)

func DriverCarSelectionRender(c *fiber.Ctx) error {
	return c.Render("carsDriver", fiber.Map{})
}
func DriverCarShow(c *fiber.Ctx) error {
	idDriver, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to convert id",
			"error":   err.Error(),
		})
	}
	var driver models.Driver
	var driversCar models.Car
	initializers.DB.First(&driver, idDriver)
	err = initializers.DB.Model(&driver).Association("Car").Find(&driversCar)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to convert id",
			"error":   err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(driversCar)
}
func DriverCreate(c *fiber.Ctx) error {
	// Получить пользователя
	var driver models.Driver
	err := c.BodyParser(&driver)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse form to driver model",
			"error":   err.Error(),
		})
	}
	// Создать водителя
	result := initializers.DB.Create(&driver)

	if result.Error != nil {

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to create a driver",
			"error":   result.Error.Error(),
		})
	}
	// Вернуть её
	return c.Status(200).JSON(fiber.Map{
		"message": "driver successfully created",
		"driver":  driver,
	})
}
func DriversIndex(c *fiber.Ctx) error {
	var drivers []models.Driver
	result := initializers.DB.Find(&drivers)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to find any drivers",
			"error":   result.Error.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "found drivers successfully",
		"drivers": drivers,
	})
}
func DriverShow(c *fiber.Ctx) error {
	id := c.Params("id")

	//Получить машину с нужным id
	var driver models.Driver
	result := initializers.DB.First(&driver, id)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to find a driver with given id",
			"error":   result.Error.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("found driver id:%v successfully", id),
		"driver":  driver,
	})
}

func DriverCarUpdate(c *fiber.Ctx) error {
	idDriver := c.Params("id")
	var driver models.Driver
	result := initializers.DB.First(&driver, idDriver)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("failed to find driver with id %v", idDriver),
			"error":   result.Error.Error(),
		})
	}

	var newCarIDRequest map[string]string
	err := c.BodyParser(&newCarIDRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse new car id",
			"error":   err.Error(),
		})
	}
	var car models.Car
	result = initializers.DB.First(&car, newCarIDRequest["carId"])
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to find new car",
			"error":   result.Error.Error(),
		})
	}
	err = initializers.DB.Model(&driver).Association("Car").Replace(&car)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to associate new car with driver",
			"error":   err,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "updated car successfully"})
}

func DriverUpdate(c *fiber.Ctx) error {
	//Получить URL с id
	id := c.Params("id")
	// Получить машину
	var driverNew models.Driver
	var driverOld models.Driver
	err := c.BodyParser(&driverNew)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse form to driver",
			"error":   err.Error(),
		})
	}
	//Получить водителя с нужным id
	result := initializers.DB.First(&driverOld, id)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("failed to find driver with id %v", id),
			"error":   result.Error.Error(),
		})
	}
	result = initializers.DB.Model(&driverOld).Omit("CarID", "Car", "Orders").Updates(&driverNew)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("failed to update driver with id %v", id),
			"error":   result.Error.Error(),
		})
	}
	var car models.Car
	result = initializers.DB.First(&car, driverNew.CarID)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to find new car",
			"error":   result.Error.Error(),
		})
	}
	errf := initializers.DB.Model(&driverOld).Association("Car").Replace(&car)
	if errf != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to associate new car with driver",
			"error":   err,
		})
	}
	// Обновить данные
	//Ответить обновленными данными
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("updated driver with id %v successfully", id),
		"driver":  driverOld,
	})
}
func DriverDelete(c *fiber.Ctx) error {
	//Получить id пользователя
	id := c.Params("id")
	//Удалить машину
	result := initializers.DB.Delete(&models.Driver{}, id)
	//Ответить о результате
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to delete driver",
			"error":   result.Error.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "driver deleted successfully",
	})

}
