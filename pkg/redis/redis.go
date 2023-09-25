package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {

	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("Failed to load env file!")
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPass := os.Getenv("REDIS_PASSWORD")

	redisURL := fmt.Sprintf("%s:%s", redisHost, redisPort)

	Rdb = redis.NewClient(&redis.Options{
		Addr:     redisURL,  // Redis server address
		Password: redisPass, // Password, leave empty if none
		DB:       0,         // Default DB
		Protocol: 3,         // Protocal specify 2 for RESP 2 or 3 for RESP 3
	})
}
