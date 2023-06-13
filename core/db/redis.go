package db

import (
	"fmt"
	"slyfox-tails/config"

	"github.com/go-redis/redis"
)

func ConnectRedis(cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})
	_, err := client.Ping().Result()

	return client, err
}
