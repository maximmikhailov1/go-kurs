package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maximmikhailov1/go-kurs/initializers"
	"github.com/maximmikhailov1/go-kurs/models"
)

func CarsCreate(c *gin.Context) {
	// Получить машину

	// Создать машину
	car := models.Cars{
		Firm_name:        "Mercedes-Benz",
		Model_name:       "Maybach S-class",
		Reg_plate_number: "MM000MMM",
		Rent:             20000,
		Is_detailed:      false,
		Is_being_used:    false}
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
