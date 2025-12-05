package main

import (
	"log"
)

func main() {
	cfg := config.Init()

	svc := service.Init(&cfg)

	hdlr := handler.Init(&cfg, svc)

	server := server.Init(cfg, hdlr)

	err := server.Router.Listen(cfg.Port)
	if err != nil {
		log.Fatalf("Error starting server", err)
	}

}
