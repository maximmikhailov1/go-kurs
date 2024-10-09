package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maximmikhailov1/go-kurs/controllers"
	"github.com/maximmikhailov1/go-kurs/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// TODO: Человек пришёл со своей машиной
func main() {
	r := gin.Default()
	r.POST("/cars", controllers.CarCreate)
	r.GET("/cars", controllers.CarsIndex)
	r.GET("/cars/:id", controllers.CarShow)
	r.PUT("/cars/:id", controllers.CarUpdate)
	r.DELETE("/cars/:id", controllers.CarDelete)
	r.POST("/users", controllers.UserCreate)
	r.GET("/users", controllers.UsersIndex)
	r.GET("/users/:id", controllers.UserShow)
	r.PUT("/users/:id", controllers.UserUpdate)
	r.DELETE("/user/:id", controllers.UserDelete)
	r.Run() // listen and serve on localhost:3000
}
