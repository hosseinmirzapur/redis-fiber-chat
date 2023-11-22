package client

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		DB:       0,
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return client
}

func RedisSet(c *redis.Client, key string, value interface{}, ttl time.Duration) error {
	return c.Set(context.Background(), key, value, ttl).Err()
}

func RedisGet(c *redis.Client, key string) (string, error) {
	val, err := c.Get(context.Background(), key).Result()

	if err == redis.Nil {
		return "", fmt.Errorf("key '%s' not found", key)
	}
	if err != nil {
		return "", err
	}

	return val, nil
}
