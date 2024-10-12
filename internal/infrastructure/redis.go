package infrastructure

import (
	"fmt"

	"github.com/didikz/goshu/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(rc config.Redis) *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rc.Host, rc.Port),
		Password: rc.Password,
		DB:       rc.Database,
	})
	return redis
}
