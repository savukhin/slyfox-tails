package api

import (
	"crypto/rsa"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func login(db *gorm.DB, privateKey *rsa.PrivateKey) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.FormValue("user")
		pass := c.FormValue("pass")

		// Throws Unauthorized error
		if user != "john" || pass != "doe" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Create the Claims
		claims := jwt.MapClaims{
			"name":  "John Doe",
			"admin": true,
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString(privateKey)
		if err != nil {
			log.Printf("token.SignedString: %v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"token": t})
	}
}
