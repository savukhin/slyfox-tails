package db

import (
	"fmt"
	"slyfox-tails/config"
	"slyfox-tails/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func autoMigrate(gormdb *gorm.DB) error {
	migrateModels := []interface{}{
		models.User{},
		models.Project{},
		models.Job{},
		models.Stage{},
		models.Point{},
	}

	return gormdb.AutoMigrate(migrateModels...)
}

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresUser,
		cfg.PostgresPassword, cfg.PostgresDB,
		cfg.PostgresPort,
	)

	postgresConn := postgres.Open(dsn)
	gormDB, err := gorm.Open(postgresConn)

	if err != nil {
		return gormDB, err
	}

	if cfg.PostgresAutoMigrate {
		err := autoMigrate(gormDB)
		if err != nil {
			return nil, err
		}
	}

	return gormDB, nil
}
