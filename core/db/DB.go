package db

import (
	"fmt"
	"slyfox-tails/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	host := utils.GetEnvDefault("POSTGRES_HOST", "localhost")
	user := utils.GetEnvDefault("POSTGRES_USER", "20624880")
	password := utils.GetEnvDefault("POSTGRES_PASSWORD", "admin")
	db := utils.GetEnvDefault("POSTGRES_DB", "slyfox-tails")
	port := utils.GetEnvDefault("POSTGRES_PORT", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, db, port,
	)

	postgresConn := postgres.Open(dsn)
	return gorm.Open(postgresConn)
}
