package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maximmikhailov1/go-kurs/initializers"
	"github.com/maximmikhailov1/go-kurs/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// TODO: Человек пришёл со своей машиной
func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	s := &http.Server{
		Addr:           os.Getenv("LOCAL_URL"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe() // listen and serve on localhost:3000
}
