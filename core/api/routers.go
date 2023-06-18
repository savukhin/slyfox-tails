package api

import (
	"crypto/rsa"
	"slyfox-tails/config"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	PointContextKey string = "point"
)

func SetupRouter(db *gorm.DB, redisClient *redis.Client, privateKey *rsa.PrivateKey, logger *zap.Logger, cfg *config.Config) *fiber.App {
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
		// Claims: jwt.MapClaims{},
		Claims: &UserClaims{},
	})

	// jwtMiddleware := JWTMiddleware(db, privateKey.Public())
	// j.MiddlewareFunc()

	// jwtPointMiddleware := jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{
	// 		JWTAlg: jwtware.RS256,
	// 		Key:    privateKey.Public(),
	// 	},
	// 	ContextKey: PointContextKey,
	// })

	user := v1.Group("/user")

	user.Get("/login", login(db, privateKey, validate))
	user.Post("/register", register(db, redisClient, logger, validate, cfg))
	user.Get("/email-verify/:code", verify(db, redisClient, logger))

	point := v1.Group("/point")

	point.Post("/login", login(db, privateKey, validate))

	v1.Get("/restricted", jwtMiddleware, restricted)

	v1.Get("users/*", NotImplemented)
	v1.Post("users/", NotImplemented)

	v1.Get("project/:project_id", jwtMiddleware, getProject(db))
	v1.Post("project/", jwtMiddleware, createProject(db, validate))
	v1.Patch("project/:project_id", jwtMiddleware, updateProject(db, validate))
	v1.Delete("project/:project_id", jwtMiddleware, deleteProject(db))

	v1.Get("job/:job_id", jwtMiddleware, getJob(db))
	v1.Post("job/", jwtMiddleware, createJob(db, validate))
	v1.Patch("job/:job_id", jwtMiddleware, updateJob(db, validate))
	v1.Delete("job/:job_id", jwtMiddleware, deleteJob(db))

	v1.Get("point/:point_id", jwtMiddleware, getPoint(db))
	v1.Post("point", jwtMiddleware, createPoint(db, validate))
	v1.Patch("point/:point_id", jwtMiddleware, updatePoint(db, validate))
	v1.Delete("point/:point_id", jwtMiddleware, deletePoint(db))

	v1.Get("stage/:stage_id", jwtMiddleware, getStage(db))
	v1.Post("stage/", jwtMiddleware, createStage(db, validate))
	v1.Patch("stage/:stage_id", jwtMiddleware, updateStage(db, validate))
	v1.Delete("stage/:stage_id", jwtMiddleware, deleteStage(db))

	return app
}
