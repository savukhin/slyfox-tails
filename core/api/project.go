package api

import (
	"slyfox-tails/db/models"
	"slyfox-tails/db/query"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func createProject(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userLocal := c.Locals("user").(*jwt.Token)
		claims := userLocal.Claims.(jwt.MapClaims)
		id := claims["id"].(int)

		payload := &CreateProjectDTO{}

		if err := c.BodyParser(payload); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		u := query.Use(db).User
		user, _ := u.Where(u.ID.Eq(id)).First()

		err := u.Projects.Model(user).Append(&models.Project{Title: payload.Title})
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func getProject(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userLocal := c.Locals("user").(*jwt.Token)
		claims := userLocal.Claims.(jwt.MapClaims)
		id := claims["id"].(int)

		payload := &CreateProjectDTO{}

		if err := c.BodyParser(payload); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		u := query.Use(db).User
		user, _ := u.Where(u.ID.Eq(id)).First()

		err := u.Projects.Model(user).Append(&models.Project{Title: payload.Title})
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
