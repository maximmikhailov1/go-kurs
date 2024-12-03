package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func ParseClientJWT(jwtTokenString string) fiber.Map {
	var data fiber.Map
	if jwtTokenString != "" {
		tokenByte, err := jwt.Parse(jwtTokenString, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			return fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)}
		}

		claims := tokenByte.Claims.(jwt.MapClaims)
		data = fiber.Map{
			"ID":         claims["id"],
			"FirstName":  claims["firstName"],
			"SecondName": claims["secondName"],
		}
	}
	return data
}
