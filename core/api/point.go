package api

import (
	"crypto/rsa"
	"fmt"
	"log"
	"slyfox-tails/db/models"
	"slyfox-tails/db/query"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

		if payload.Password != payload.PasswordRepeat {
			return fiber.NewError(fiber.ErrBadRequest.Code, "Passwords doesn't match")
		}

		userLocal := c.Locals("user").(*jwt.Token)
		claims, ok := userLocal.Claims.(*UserClaims)
		if !ok {
			return fiber.ErrInternalServerError
		}

		userID := claims.UserID

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		s := query.Use(db).Stage
		stages := make([]*models.Stage, len(payload.Stages))
		for i, stageID := range payload.Stages {
			stage, err := s.Where(s.CreatorID.Eq(userID), s.ID.Eq(stageID)).First()
			if err != nil {
				c.Status(fiber.StatusNotFound)
				return c.JSON(map[string]string{"message": "stage id " + strconv.FormatUint(stageID, 10) + " Not exists"})
			}

			stages[i] = stage
		}

		pt := query.Use(db).Point
		point := &models.Point{
			CreatorID:    userID,
			Title:        payload.Title,
			Login:        payload.Login,
			PasswordHash: string(hashedPassword),
		}

		existing, err := pt.
			Where(
				pt.CreatorID.Eq(userID),
				pt.Login.Eq(point.Login),
				pt.DeletedAt.Null(),
			).Take()

		fmt.Println("existing ", existing, err)

		if existing != nil {
			return fiber.ErrConflict
		}

		err = pt.Create(point)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.SendString(err.Error())
		}

		fmt.Println(stages)

		err = pt.Stages.Model(point).Append(stages...)
		if err != nil {
			return fiber.ErrInternalServerError
		}

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

func loginPoint(db *gorm.DB, privateKey *rsa.PrivateKey, validate validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := &LoginPointDTO{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		if err := validate.Struct(payload); err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, "Incorrect data")
		}

		splitted := strings.Split(payload.Username, "@")
		if len(splitted) != 2 {
			return fiber.ErrBadRequest
		}

		pointLogin := splitted[0]
		username := splitted[1]

		pt := query.Use(db).Point

		u := query.Use(db).User

		point, err := pt.
			Join(u, pt.CreatorID.EqCol(u.ID)).
			Where(
				u.Username.Eq(username),
				u.EmailVerified.Is(true),
				pt.Login.Eq(pointLogin),
			).
			First()

		fmt.Println("err?", err)
		// Throws Unauthorized error
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if err := bcrypt.CompareHashAndPassword([]byte(point.PasswordHash), []byte(payload.Password)); err != nil {
			fmt.Println("pass?", err)
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		claims := PointClaims{
			PointID:   point.ID,
			ExpiresAt: time.Now().Add(time.Hour * 72),
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

func restrictedPoint(c *fiber.Ctx) error {
	user := c.Locals("point").(*jwt.Token)
	claims := user.Claims.(*PointClaims)
	fmt.Println("claim is", claims)
	userid := claims.PointID
	fmt.Println("userid is", userid)
	return c.SendString("Welcome " + strconv.FormatUint(userid, 10))
}
