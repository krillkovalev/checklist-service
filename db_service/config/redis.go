package config

import (
	"context"
	"github.com/redis/go-redis/v9"
    "os"
    "log"
	"github.com/joho/godotenv"
)

func RedisConnection(ctx context.Context) *redis.Client{
    err := godotenv.Load(".env")
	if err != nil {
		log.Println("Failed to load .env file")
	}
    address := os.Getenv("REDIS_ADDR")
	client := redis.NewClient(&redis.Options{
        Addr:     address,
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    // Ping Redis to check if the connection is working
    _, err = client.Ping(ctx).Result()
    if err != nil {
        panic(err) 
    }

	return client
}