package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"slyfox-tails/api"
	"slyfox-tails/config"
	"slyfox-tails/db"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

func main() {
	cfg := &config.Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		panic(err)
	}

	redisClient, err := db.ConnectRedis(cfg)
	if err != nil {
		panic(err)
	}

	var logger *zap.Logger
	if cfg.Mode == config.ReleaseMode {
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

	app := api.SetupRouter(gormDB, redisClient, privateKey, logger, cfg)

	app.Listen(cfg.Port)
}
