package main

import (
	"fmt"

	"github.com/didikz/goshu/config"
	"github.com/didikz/goshu/internal"
	"github.com/didikz/goshu/internal/infrastructure"
)

func main() {
	cfg := config.Load("./")
	port := fmt.Sprintf(":%d", cfg.App.Port)

	db := infrastructure.NewDB(cfg.Database)
	redis := infrastructure.NewRedisClient(cfg.Redis)

	server := internal.NewServer(port, db, redis)
	server.Run()
}
