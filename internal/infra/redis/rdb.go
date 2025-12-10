package redis

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func NewrdbClient() *redis.Client {
	rdbAddress := os.Getenv("REDIS_ADDRESS")
	redbPass := os.Getenv("REDIS_PASS")
	rdb := redis.NewClient(&redis.Options{
		Addr:     rdbAddress,
		Password: redbPass, // optional
		DB:       0,
	})
	return rdb
}
