package api

import (
	"fmt"
	"slyfox-tails/db/models"
	"slyfox-tails/db/query"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func createPoint(db *gorm.DB, validate *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := &CreatePointDTO{}

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

		s := query.Use(db).Stage
		fmt.Println("IDS = ", payload.Stages)
		stages := make([]*models.Stage, len(payload.Stages))
		for i, stageID := range payload.Stages {
			stage, err := s.Where(s.CreatorID.Eq(userID), s.ID.Eq(stageID)).First()
			if err != nil {
				c.Status(fiber.StatusNotFound)
				return c.JSON(map[string]string{"message": "stage id " + strconv.FormatUint(stageID, 10) + " Not exists"})
			}

			// stages = append(stages, stage)
			stages[i] = stage
		}

		pt := query.Use(db).Point
		point := &models.Point{CreatorID: userID, Title: payload.Title}
		err := pt.Create(point)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.SendString(err.Error())
		}

		fmt.Println(stages)

		// err = pt.Stages.Model(point).
		err = pt.Stages.Model(point).Append(stages...)
		fmt.Println("Inserted?")
		if err != nil {
			return fiber.ErrInternalServerError
		}

		// point, _ = pt.Where(pt.ID.Eq(point.ID)).First()

		c.Status(fiber.StatusCreated)
		return c.JSON(map[string]uint64{"id": point.ID})
	}
}

func getPoint(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userLocal := c.Locals("user").(*jwt.Token)
		claims, ok := userLocal.Claims.(*UserClaims)
		if !ok {
			return fiber.ErrInternalServerError
		}
		userID := claims.UserID
		pointID, err := strconv.ParseUint(c.Params("point_id"), 10, 64)
		if err != nil {
			return fiber.ErrBadRequest
		}

		pt := query.Use(db).Point
		point, err := pt.Where(pt.CreatorID.Eq(userID), pt.ID.Eq(pointID)).First()
		if err != nil {
			return fiber.ErrForbidden
		}

		stages, err := pt.Stages.Model(point).Find()

		if err == nil {
			point.Stages = stages
		}

		return c.JSON(point)
	}
}

func updatePoint(db *gorm.DB, validate *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := &UpdatePointDTO{}

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
		pointID, err := strconv.ParseUint(c.Params("point_id"), 10, 64)
		if err != nil {
			return fiber.ErrBadRequest
		}

		pt := query.Use(db).Point
		point, err := pt.Where(pt.ID.Eq(pointID), pt.CreatorID.Eq(userID)).First()

		if err != nil {
			return fiber.ErrForbidden
		}

		if payload.Stages != nil {
			s := query.Use(db).Stage
			stages := make([]*models.Stage, len(payload.Stages))
			for _, stageID := range payload.Stages {
				stage, err := s.Where(s.CreatorID.Eq(userID), s.ID.Eq(stageID)).First()
				if err != nil {
					c.Status(fiber.StatusNotFound)
					return c.JSON(map[string]string{"message": "stage id " + strconv.FormatUint(stageID, 10) + " Not exists"})
				}

				stages = append(stages, stage)
			}

			point.Stages = stages
		}

		if payload.Title != "" {
			point.Title = payload.Title
		}

		err = pt.Save(point)

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.SendString(err.Error())
		}

		stages, err := pt.Stages.Model(point).Find()

		if err == nil {
			point.Stages = stages
		}

		c.Status(fiber.StatusAccepted)
		return c.JSON(point)
	}
}

func deletePoint(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userLocal := c.Locals("user").(*jwt.Token)
		claims, ok := userLocal.Claims.(*UserClaims)
		if !ok {
			return fiber.ErrInternalServerError
		}
		userID := claims.UserID

		pointID, err := strconv.ParseUint(c.Params("point_id"), 10, 64)
		if err != nil {
			return fiber.ErrBadRequest
		}

		pt := query.Use(db).Point
		res, err := pt.Where(pt.CreatorID.Eq(userID), pt.ID.Eq(pointID)).Delete()
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

func LoginPoint() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}
