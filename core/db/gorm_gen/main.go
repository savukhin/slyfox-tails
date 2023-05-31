package main

import (
	"slyfox-tails/db/models"

	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// dsn := "host=localhost user=20624880 password=admin dbname=db_08 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// gormdb, _ := gorm.Open(postgres.Open(dsn))
	// g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(
		models.User{},
		models.Project{},
		models.Job{},
		models.Stage{},
		models.Point{},
	)

	// Generate the code
	g.Execute()
}
