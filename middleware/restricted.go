package middleware

import (
	"github.com/gofiber/fiber/v2"
)

var routes = []string{"/order-taxi", ""}

func Authorize(c *fiber.Ctx) error {
	//fmt.Println(string(c.Body()))
	path := c.Path()
	//fmt.Println(path)
	for _, v := range routes {
		if path == v && c.Cookies("authClient") == "" {
			return c.Redirect("/authClient")
		}
	}
	return c.Next()
}
