package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(host, pwd string, port int, db int) *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: pwd,
		DB:       db,
	})
	return redis
}
