package main

import (
	"slyfox-tails/api"
	"slyfox-tails/utils"
)

func main() {
	PORT := utils.GetEnvDefault("PORT", ":8080")
	// MODE := utils.GetEnvDefault("MODE", "release")

	app := api.SetupRouter()

	app.Listen(PORT)
}
