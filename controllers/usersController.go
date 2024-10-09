package controllers

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/maximmikhailov1/go-kurs/initializers"
	"github.com/maximmikhailov1/go-kurs/models"
	"gorm.io/gorm"
)

func UserCreate(c *gin.Context) {
	// Получить пользователя
	var body struct {
		First_name string
		Last_name  string
		Patronymic string
		Username   string
		CarID      uint
	}
	c.Bind(&body)
	// Создать пользователя
	user := models.User{
		First_name: body.First_name,
		Last_name:  body.Last_name,
		Patronymic: body.Patronymic,
		Username:   body.Username,
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			c.JSON(400, gin.H{"message": "Username already exists"})
			log.Println(result.Error)
		}
		c.Status(400)

		return
	}
	// Вернуть её
	c.JSON(200, gin.H{
		"user": user,
	})
}
func UsersIndex(c *gin.Context) {
	var users []models.User
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}
func UserShow(c *gin.Context) {
	//Получить URL с id
	id := c.Param("id")

	//Получить пользователя с нужным id
	var user models.User
	result := initializers.DB.First(&user, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
func UserUpdate(c *gin.Context) {
	//Получить URL с id
	id := c.Param("id")

	// Получить пользователя
	var body struct {
		First_name string
		Last_name  string
		Patronymic string
		Username   string
		CarID      uint
	}
	c.Bind(&body)
	//Получить машину с нужным id
	var user models.User
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		c.Status(400)
		return
	}
	// Обновить данные
	initializers.DB.Model(&user).Updates(models.User{
		First_name: body.First_name,
		Last_name:  body.Last_name,
		Patronymic: body.Patronymic,
		Username:   body.Username,
		CarID:      body.CarID,
	})
	//Ответить обновленными данными
	c.JSON(200, gin.H{
		"user": user,
	})
}
func UserDelete(c *gin.Context) {
	//Получить id пользователя
	id := c.Param("id")
	//Удалить машину
	result := initializers.DB.Delete(&models.User{}, id)
	//Ответить о результате
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.Status(200)

}
