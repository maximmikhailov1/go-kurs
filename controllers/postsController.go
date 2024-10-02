package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maximmikhailov1/go-kurs/initializers"
	"github.com/maximmikhailov1/go-kurs/models"
)

func CarsCreate(c *gin.Context) {
	// Получить машину
	var body struct {
		Firm_name        string
		Model_name       string
		Reg_plate_number string
		Rent             int
		Is_detailed      bool
		Is_being_used    bool
	}
	c.Bind(&body)
	// Создать машину
	car := models.Cars{
		Firm_name:        body.Firm_name,
		Model_name:       body.Model_name,
		Reg_plate_number: body.Reg_plate_number,
		Rent:             body.Rent,
		Is_detailed:      body.Is_detailed,
		Is_being_used:    body.Is_detailed}
	result := initializers.DB.Create(&car)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Вернуть её
	c.JSON(200, gin.H{
		"car": car,
	})
}
func CarsIndex(c *gin.Context) {
	var cars []models.Cars
	result := initializers.DB.Find(&cars)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"cars": cars,
	})
}
func CarsShow(c *gin.Context) {
	//Получить URL с id
	id := c.Param("id")

	//Получить машину с нужным id
	var car models.Cars
	result := initializers.DB.First(&car, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"car": car,
	})
}
func CarsUpdate(c *gin.Context) {
	//Получить URL с id
	id := c.Param("id")

	// Получить машину
	var body struct {
		Firm_name        string
		Model_name       string
		Reg_plate_number string
		Rent             int
		Is_detailed      bool
		Is_being_used    bool
	}
	c.Bind(&body)
	//Получить машину с нужным id
	var car models.Cars
	result := initializers.DB.First(&car, id)
	if result.Error != nil {
		c.Status(400)
		return
	}
	// Обновить данные
	initializers.DB.Model(&car).Updates(models.Cars{
		Firm_name:        body.Firm_name,
		Model_name:       body.Model_name,
		Reg_plate_number: body.Reg_plate_number,
		Rent:             body.Rent,
		Is_detailed:      body.Is_detailed,
		Is_being_used:    body.Is_being_used,
	})
	//Ответить обновленными данными
	c.JSON(200, gin.H{
		"car": car,
	})
}
func CarsDelete(c *gin.Context) {
	//Получить id машины
	id := c.Param("id")
	//Удалить машину
	result := initializers.DB.Delete(&models.Cars{}, id)
	//Ответить о результате
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.Status(200)

}
