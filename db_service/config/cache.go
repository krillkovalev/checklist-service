package config

import (
	"github.com/go-redis/redis"
)

func RedisConnection() *redis.Client{

	client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    // Ping Redis to check if the connection is working
    _, err := client.Ping().Result()
    if err != nil {
        panic(err)
    }

	return client
}