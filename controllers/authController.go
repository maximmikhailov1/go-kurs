package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/maximmikhailov1/go-kurs/initializers"
	"github.com/maximmikhailov1/go-kurs/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func AuthRender(c *fiber.Ctx) error {
	return c.Render("authClient", fiber.Map{})
}

func SignInClient(c *fiber.Ctx) error {
	timeStart := time.Now().UnixMilli()
	var clientForm models.Client
	var client models.Client
	err := c.BodyParser(&clientForm)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse a client form to sign in an account",
			"error":   err.Error(),
		})
	}
	fmt.Println("TO PARSE", time.Now().UnixMilli()-timeStart)
	timeStart = time.Now().UnixMilli()
	result := initializers.DB.First(&client, "username = ?", clientForm.Username)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid username or password",
		})
	}
	fmt.Println("TO FIND", time.Now().UnixMilli()-timeStart)
	timeStart = time.Now().UnixMilli()
	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(clientForm.Password))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid username or password",
		})
	}
	fmt.Println("TO HASH", time.Now().UnixMilli()-timeStart)
	timeStart = time.Now().UnixMilli()
	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now()
	claims := tokenByte.Claims.(jwt.MapClaims)
	claims["id"] = client.ID
	claims["firstName"] = client.FirstName
	claims["secondName"] = client.SecondName
	claims["exp"] = now.Add(time.Hour * 24 * 30).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": fmt.Sprintf("generating JWT Token failed: %v", err)})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "authClient",
		Value:   tokenString,
		Path:    "/",
		Expires: now.Add(time.Hour * 24 * 30),
		Secure:  false,
	})

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully logged in",
	})
}

func SignUpClient(c *fiber.Ctx) error {
	var client models.Client

	err := c.BodyParser(&client)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse a client form to create an account",
			"error":   err.Error(),
		})
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(client.Password), 16)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to hash the password",
		})
	}
	client.Password = string(hash)

	result := initializers.DB.Create(&client)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to create a client account",
			"error":   result.Error.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "client account successfully created",
	})
}

func SignInDriver(c *fiber.Ctx) error {
	timeStart := time.Now().UnixMilli()
	var driverForm models.Driver
	var driver models.Driver
	err := c.BodyParser(&driverForm)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse a client form to sign in an account",
			"error":   err.Error(),
		})
	}
	fmt.Println("TO PARSE", time.Now().UnixMilli()-timeStart)
	timeStart = time.Now().UnixMilli()
	result := initializers.DB.First(&driver, "username = ?", driverForm.Username)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid username or password",
		})
	}
	fmt.Println("TO FIND", time.Now().UnixMilli()-timeStart)
	timeStart = time.Now().UnixMilli()
	err = bcrypt.CompareHashAndPassword([]byte(driver.Password), []byte(driverForm.Password))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid username or password",
		})
	}
	fmt.Println("TO HASH", time.Now().UnixMilli()-timeStart)
	timeStart = time.Now().UnixMilli()
	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now()
	claims := tokenByte.Claims.(jwt.MapClaims)
	claims["id"] = driver.ID
	claims["firstName"] = driver.FirstName
	claims["secondName"] = driver.SecondName
	claims["exp"] = now.Add(time.Hour * 24 * 30).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": fmt.Sprintf("generating JWT Token failed: %v", err)})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "authDriver",
		Value:   tokenString,
		Path:    "/",
		Expires: now.Add(time.Hour * 24 * 30),
		Secure:  false,
	})

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully logged in",
	})
}

func SignUpDriver(c *fiber.Ctx) error {
	var driver models.Driver

	err := c.BodyParser(&driver)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse a client form to create an account",
			"error":   err.Error(),
		})
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(driver.Password), 16)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to hash the password",
		})
	}
	driver.Password = string(hash)

	result := initializers.DB.Create(&driver)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to create a client account",
			"error":   result.Error.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "client account successfully created",
	})
}

func Logout(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24 * 30)
	c.Cookie(&fiber.Cookie{
		Name:    "authClient",
		Value:   "",
		Expires: expired,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "authDriver",
		Value:   "",
		Expires: expired,
	})
	return c.Redirect("/")
}
