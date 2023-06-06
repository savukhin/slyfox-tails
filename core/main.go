package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"slyfox-tails/api"
	"slyfox-tails/db"
	"slyfox-tails/utils"

	"go.uber.org/zap"
)

const (
	releaseMode string = "release"
	debugMode   string = "debug"
	testMode    string = "release"
)

func main() {
	MODE := utils.GetEnvDefault("MODE", "release")
	PORT := utils.GetEnvDefault("PORT", ":8080")

	gormDB, err := db.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	redisClient, err := db.ConnectRedis()
	if err != nil {
		panic(err)
	}

	var logger *zap.Logger
	if MODE == "release" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugaredLogger := logger.Sugar()
	defer sugaredLogger.Sync()

	rng := rand.Reader
	privateKey, err := rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}

	app := api.SetupRouter(gormDB, redisClient, privateKey, logger)

	app.Listen(PORT)
}
