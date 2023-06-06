package api

import (
	"crypto/rsa"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, redisClient *redis.Client, privateKey *rsa.PrivateKey, logger *zap.Logger) *fiber.App {
	app := fiber.New()

	validate := validator.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// 	// AllowMethods:     "GET,POST,PATCH,DELETE",
	// 	AllowCredentials: false,
	// }))

	app.Get("/ping", Pong)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    privateKey.Public(),
		},
	})

	v1.Post("/login", login(db, privateKey, validate))
	v1.Post("/register", register(db, redisClient, logger, validate))
	v1.Post("/email-verify/:code", verify(db, redisClient, logger))

	v1.Get("/restricted", jwtMiddleware, func(c *fiber.Ctx) error {
		fmt.Println(c.Locals("user"))
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		return c.SendString("Welcome " + name)
	})

	v1.Get("users/*", NotImplemented)
	v1.Post("users/", NotImplemented)

	v1.Get("project/*", NotImplemented)
	v1.Post("project/*", NotImplemented)
	v1.Patch("project/*", NotImplemented)
	v1.Delete("project/*", NotImplemented)

	v1.Get("job/*", NotImplemented)
	v1.Post("job/*", NotImplemented)
	v1.Patch("job/*", NotImplemented)
	v1.Delete("job/*", NotImplemented)

	v1.Get("point/*", NotImplemented)
	v1.Post("point/*", NotImplemented)
	v1.Patch("point/*", NotImplemented)
	v1.Delete("point/*", NotImplemented)

	v1.Get("stage/*", NotImplemented)
	v1.Post("stage/*", NotImplemented)
	v1.Patch("stage/*", NotImplemented)
	v1.Delete("stage/*", NotImplemented)

	return app
}
