package main

import (
	"log"

	"kolesagpt/config"
	"kolesagpt/internal/handler"
	"kolesagpt/pkg/server"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	handler := handler.NewHandler(cfg)
	if err := server.Serve(cfg.Port, handler.Routes()); err != nil {
		log.Fatal(err)
	}
}
