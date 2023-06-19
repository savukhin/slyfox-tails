package api

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"slyfox-tails/config"
	"slyfox-tails/db/models"
	"slyfox-tails/db/query"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	verificationCodeTTL = 24 * time.Hour
)

func login(db *gorm.DB, privateKey *rsa.PrivateKey, validate *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := &LoginUserDTO{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		if err := validate.Struct(payload); err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, "Incorrect data")
		}

		u := query.Use(db).User
		user, err := u.
			Where(u.Username.Eq(payload.Username), u.EmailVerified.Is(true)).
			First()

		// Throws Unauthorized error
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(payload.Password)); err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Create the Claims
		// claims := jwt.MapClaims{
		// 	"user_id": user.ID,
		// 	"admin":   true,
		// 	"exp":     time.Now().Add(time.Hour * 72).Unix(),
		// }

		claims := UserClaims{
			UserID:    user.ID,
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

func register(db *gorm.DB, redisClient *redis.Client, logger *zap.Logger, validate *validator.Validate, cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("Register")
		payload := &RegisterUserDTO{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		if err := validate.Struct(payload); err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, "Incorrect data")
		}

		if payload.Password != payload.PasswordRepeat {
			return fiber.NewError(fiber.ErrBadRequest.Code, "Passwords doesn't match")
		}

		u := query.Use(db).User
		_, err := u.Where(u.Username.Eq(payload.Username)).Or(u.Email.Eq(payload.Email)).First()
		if err == nil {
			return fiber.NewError(fiber.ErrConflict.Code, "Username of email already exists")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		userModel := &models.User{
			Username:      payload.Username,
			Email:         payload.Email,
			PasswordHash:  string(hashedPassword),
			EmailVerified: (cfg.Mode == config.TestMode),
		}

		if err := u.Create(userModel); err != nil {
			return err
		}

		if cfg.Mode != config.TestMode {
			hasher := sha256.New()
			hasher.Write([]byte(time.Now().String()))
			verificationCode := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

			logger.Info(fmt.Sprintf("Generated verificationCode '%s'", verificationCode))

			redisClient.Set(verificationCode, userModel.ID, verificationCodeTTL)
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func verify(db *gorm.DB, redisClient *redis.Client, logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		verificactionCode := c.Params("code")
		userIDStr, err := redisClient.Get(verificactionCode).Result()

		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		userID, err := strconv.ParseUint(userIDStr, 10, 64)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		u := query.Use(db).User
		result, err := u.
			Where(u.ID.Eq(userID)).
			Update(u.EmailVerified, true)

		if err != nil || result.RowsAffected != 1 {
			return c.SendStatus(fiber.StatusNotFound)
		}

		logger.Info(fmt.Sprintf("User %d verified", userID))

		return c.SendStatus(fiber.StatusOK)
	}
}

func restrictedUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	claims := user.Claims.(*UserClaims)
	// id := claims["id"].(float64)
	// fmt.Println("id is", id)
	// userid := claims["user_id"].(uint64)
	// fmt.Println("userid is", userid)
	fmt.Println("claim is", claims)
	userid := claims.UserID
	fmt.Println("userid is", userid)
	return c.SendString("Welcome " + strconv.FormatUint(userid, 10))
}
