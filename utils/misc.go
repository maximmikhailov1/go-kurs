package utils

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func DriverGetId(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(ParseJWT(c.Cookies("authDriver"))["ID"])
}
