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

func createJob(db *gorm.DB, validate *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := &CreateJobDTO{}

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

		p := query.Use(db).Project
		proj, err := p.Where(p.CreatorID.Eq(userID), p.ID.Eq(payload.ProjectID)).First()

		if err != nil {
			return fiber.ErrForbidden
		}

		job := &models.Job{Title: payload.Title, CreatorID: userID}
		err = p.Jobs.Model(proj).Append(job)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		c.Status(fiber.StatusCreated)
		return c.JSON(map[string]uint64{"id": job.ID})
	}
}

func getJob(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userLocal := c.Locals("user").(*jwt.Token)
		claims, ok := userLocal.Claims.(*UserClaims)
		if !ok {
			return fiber.ErrInternalServerError
		}
		userID := claims.UserID
		jobID, err := strconv.ParseUint(c.Params("job_id"), 10, 64)
		if err != nil {
			return fiber.ErrBadRequest
		}

		j := query.Use(db).Job
		job, err := j.Where(j.CreatorID.Eq(userID), j.ID.Eq(jobID)).First()

		if err != nil {
			return fiber.ErrForbidden
		}

		return c.JSON(job)
	}
}

func updateJob(db *gorm.DB, validate *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := &UpdateJobDTO{}

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

		jobID, err := strconv.ParseUint(c.Params("job_id"), 10, 64)
		if err != nil {
			return fiber.ErrBadRequest
		}

		j := query.Use(db).Job
		job, err := j.Where(j.ID.Eq(jobID), j.CreatorID.Eq(userID)).First()

		if err != nil {
			return fiber.ErrForbidden
		}

		job.Title = payload.Title
		err = j.Save(job)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		job, err = j.Where(j.CreatorID.Eq(userID), j.ID.Eq(jobID)).First()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		c.Status(fiber.StatusAccepted)
		return c.JSON(job)
	}
}

func deleteJob(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userLocal := c.Locals("user").(*jwt.Token)
		claims, ok := userLocal.Claims.(*UserClaims)
		if !ok {
			return fiber.ErrInternalServerError
		}
		userID := claims.UserID

		jobID, err := strconv.ParseUint(c.Params("job_id"), 10, 64)
		if err != nil {
			return fiber.ErrBadRequest
		}

		p := query.Use(db).Job
		res, err := p.Where(p.CreatorID.Eq(userID), p.ID.Eq(jobID)).Delete()
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
