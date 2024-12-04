package middleware

import "github.com/gofiber/fiber/v2"

func IsAuthenticated(c *fiber.Ctx) error {
	if c.Cookies("authClient") != "" || c.Cookies("authDriver") != "" {
		c.Locals("loggedIn", true)
	}
	return c.Next()
}
