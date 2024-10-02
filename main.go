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

func main() {
	r := gin.Default()
	r.POST("/cars", controllers.CarsCreate)
	r.Run() // listen and serve on localhost:8080
}
