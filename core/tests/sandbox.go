// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"slyfox-tails/db/query"
	"slyfox-tails/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	host := utils.GetEnvDefault("POSTGRES_HOST", "localhost")
	user := utils.GetEnvDefault("POSTGRES_USER", "20624880")
	password := utils.GetEnvDefault("POSTGRES_PASSWORD", "admin")
	db := utils.GetEnvDefault("POSTGRES_DB", "slyfox-tails")
	port := utils.GetEnvDefault("POSTGRES_PORT", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, db, port,
	)

	postgresConn := postgres.Open(dsn)
	gormDB, _ := gorm.Open(postgresConn)

	u := query.Use(gormDB).User
	// where := u.Where(u.ID.Eq(5))
	// mike, _ := u.Where(u.ID.Eq(5)).Update(u.Username, "MikeIn")
	// fmt.Println(mike)

	mike, _ := u.Where(u.ID.Eq(5)).Find()
	fmt.Println(mike[0].Username)
}
