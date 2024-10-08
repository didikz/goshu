package main

import (
	"fmt"
	"log"

	"github.com/didikz/goshu/config"
)

func main() {
	cfg := config.Load("./")
	port := fmt.Sprintf(":%d", cfg.Port)
	log.Println(cfg)

	db := NewDB(cfg.Database)

	server := NewServer(port, db)
	server.Run()
}
