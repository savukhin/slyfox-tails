package api

import (
	"slyfox-tails/db/models"
	"slyfox-tails/db/query"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func createProject(db *gorm.DB, validate *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := &CreateProjectDTO{}

		if err := c.BodyParser(payload); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if err := validate.Struct(payload); err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, "Incorrect data")
		}

		userLocal := c.Locals("user").(*jwt.Token)
		claims, ok := userLocal.Claims.(*UserClaims)
		if !ok {
			return fiber.ErrInternalServerError
		}

		id := claims.UserID

		u := query.Use(db).User
		user, _ := u.Where(u.ID.Eq(id)).First()

		proj := &models.Project{Title: payload.Title}
		err := u.Projects.Model(user).Append(proj)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		c.Status(fiber.StatusCreated)
		return c.JSON(map[string]uint64{"id": proj.ID})
	}
}

func getProject(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userLocal := c.Locals("user").(*jwt.Token)
		claims, ok := userLocal.Claims.(*UserClaims)
		if !ok {
			return fiber.ErrInternalServerError
		}
		userID := claims.UserID
		projectID, err := strconv.ParseUint(c.Params("project_id"), 10, 64)
		if err != nil {
			return fiber.ErrBadRequest
		}

		p := query.Use(db).Project
		proj, err := p.Where(p.CreatorID.Eq(userID), p.ID.Eq(projectID)).First()

		if err != nil {
			return fiber.ErrForbidden
		}

		return c.JSON(proj)
	}
}

func updateProject(db *gorm.DB, validate *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := &CreateProjectDTO{}

		if err := c.BodyParser(payload); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if err := validate.Struct(payload); err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, "Incorrect data")
		}

		userLocal := c.Locals("user").(*jwt.Token)
		claims, ok := userLocal.Claims.(*UserClaims)
		if !ok {
			return fiber.ErrInternalServerError
		}
		userID := claims.UserID

		projectID, err := strconv.ParseUint(c.Params("project_id"), 10, 64)
		if err != nil {
			return fiber.ErrBadRequest
		}

		project := &models.Project{
			ID:        projectID,
			Title:     payload.Title,
			CreatorID: userID,
		}

		p := query.Use(db).Project
		err = p.Save(project)

		if err != nil {
			return fiber.ErrForbidden
		}

		proj, err := p.Where(p.CreatorID.Eq(userID), p.ID.Eq(projectID)).First()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		c.Status(fiber.StatusAccepted)
		return c.JSON(proj)
	}
}

func deleteProject(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userLocal := c.Locals("user").(*jwt.Token)
		claims, ok := userLocal.Claims.(*UserClaims)
		if !ok {
			return fiber.ErrInternalServerError
		}
		userID := claims.UserID

		projectID, err := strconv.ParseUint(c.Params("project_id"), 10, 64)
		if err != nil {
			return fiber.ErrBadRequest
		}

		p := query.Use(db).Project
		res, err := p.Where(p.CreatorID.Eq(userID), p.ID.Eq(projectID)).Delete()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		if res.RowsAffected == 0 {
			return fiber.ErrNotFound
		}

		if res.RowsAffected > 1 {
			return fiber.ErrInsufficientStorage
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
