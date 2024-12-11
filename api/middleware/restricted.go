package middleware

import (
	"github.com/gofiber/fiber/v2"
)

var routesClientRestricted = []string{"/order-taxi"}
var routesDriverRestricted = []string{"/drivers/car"}

func Authorize(c *fiber.Ctx) error {
	//fmt.Println(string(c.Body()))
	path := c.Path()
	//fmt.Println(path)
	for _, v := range routesClientRestricted {
		if path == v && c.Cookies("authClient") == "" {
			return c.Redirect("/auth/client")
		}
	}
	for _, v := range routesDriverRestricted {
		if path == v && c.Cookies("authDriver") == "" {
			return c.Redirect("/auth/driver")
		}
	}
	return c.Next()
}
