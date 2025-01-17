package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func RedisConnection(ctx context.Context) *redis.Client{

	client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    // Ping Redis to check if the connection is working
    _, err := client.Ping(ctx).Result()
    if err != nil {
        panic(err)
    }

	return client
}