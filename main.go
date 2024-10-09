package main

import (
	"fmt"

	"github.com/didikz/goshu/config"
)

func main() {
	cfg := config.Load("./")
	port := fmt.Sprintf(":%d", cfg.Port)

	db := NewDB(cfg.Database)
	redis := NewRedisClient("localhost", "", 6379, 0)

	server := NewServer(port, db, redis)
	server.Run()
}
