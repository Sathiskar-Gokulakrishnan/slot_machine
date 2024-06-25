package config

import (
	"os"
)

var (
	MongoDBUri = getEnv("MONGODB_URI", "mongodb://localhost:27017")
	RedisAddr  = getEnv("REDIS_ADDR", "localhost:6379")
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
