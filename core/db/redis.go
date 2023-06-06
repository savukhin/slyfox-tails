package db

import (
	"fmt"
	"slyfox-tails/utils"

	"github.com/go-redis/redis"
)

func ConnectRedis() (*redis.Client, error) {
	host := utils.GetEnvDefault("REDIS_HOST", "localhost")
	port := utils.GetEnvDefault("REDIS_PORT", "6379")
	password := utils.GetEnvDefault("REDIS_PASSWORD", "")
	db := utils.GetEnvDefaultInt("REDIS_DB", 0)

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})
	_, err := client.Ping().Result()

	return client, err
}
