package api

import (
	"slyfox-tails/db/models"
	"slyfox-tails/db/query"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func createStage(db *gorm.DB, validate *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := &CreateStageDTO{}

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

		j := query.Use(db).Job
		job, err := j.Where(j.CreatorID.Eq(userID), j.ID.Eq(payload.JobID)).First()

		if err != nil {
			return fiber.ErrForbidden
		}

		stage := &models.Stage{Title: payload.Title, CreatorID: userID}
		err = j.Stages.Model(job).Append(stage)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		c.Status(fiber.StatusCreated)
		return c.JSON(map[string]uint64{"id": stage.ID})
	}
}

func getStage(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userLocal := c.Locals("user").(*jwt.Token)
		claims, ok := userLocal.Claims.(*UserClaims)
		if !ok {
			return fiber.ErrInternalServerError
		}
		userID := claims.UserID
		stageID, err := strconv.ParseUint(c.Params("stage_id"), 10, 64)
		if err != nil {
			return fiber.ErrBadRequest
		}

		st := query.Use(db).Stage
		stage, err := st.Where(st.CreatorID.Eq(userID), st.ID.Eq(stageID)).First()

		if err != nil {
			return fiber.ErrForbidden
		}

		return c.JSON(stage)
	}
}

func updateStage(db *gorm.DB, validate *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := &UpdatedStageDTO{}

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

		stageID, err := strconv.ParseUint(c.Params("stage_id"), 10, 64)
		if err != nil {
			return fiber.ErrBadRequest
		}

		s := query.Use(db).Stage
		stage, err := s.Where(s.ID.Eq(stageID), s.CreatorID.Eq(userID)).First()

		if err != nil {
			return fiber.ErrForbidden
		}

		stage.Title = payload.Title
		stage.StartedAt = time.UnixMilli(int64(payload.StartedAtMs))
		err = s.Save(stage)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		stage, err = s.Where(s.CreatorID.Eq(userID), s.ID.Eq(stageID)).First()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		c.Status(fiber.StatusAccepted)
		return c.JSON(stage)
	}
}

func deleteStage(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userLocal := c.Locals("user").(*jwt.Token)
		claims, ok := userLocal.Claims.(*UserClaims)
		if !ok {
			return fiber.ErrInternalServerError
		}
		userID := claims.UserID

		stageID, err := strconv.ParseUint(c.Params("stage_id"), 10, 64)
		if err != nil {
			return fiber.ErrBadRequest
		}

		s := query.Use(db).Stage
		res, err := s.Where(s.CreatorID.Eq(userID), s.ID.Eq(stageID)).Delete()
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
