package utils

import (
	"context"
	"slot-machine-api/config"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: config.RedisAddr,
	})
}

func GetRedisClient() *redis.Client {
	return RedisClient
}

var Ctx = context.Background()
