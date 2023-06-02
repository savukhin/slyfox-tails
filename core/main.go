package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"slyfox-tails/api"
	"slyfox-tails/db"
	"slyfox-tails/utils"
)

func main() {
	gormDB, err := db.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	PORT := utils.GetEnvDefault("PORT", ":8080")
	// MODE := utils.GetEnvDefault("MODE", "release")

	rng := rand.Reader
	privateKey, err := rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}

	app := api.SetupRouter(gormDB, privateKey)

	app.Listen(PORT)
}
