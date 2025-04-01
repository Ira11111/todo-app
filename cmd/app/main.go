package main

import (
	"github.com/Ira11111/todo-app/internal/config"
	"github.com/Ira11111/todo-app/internal/handlers"
	"github.com/Ira11111/todo-app/internal/transport"
	"log"
)

func main() {
	cfg := config.MustLoad()
	appHandlers := &handlers.Handler{}
	server := new(transport.Server)

	if err := server.Start(
		cfg.Server.Port,
		cfg.Server.Timeout,
		cfg.Server.IdleTimeout,
		appHandlers.InitRoutes(),
	); err != nil {
		log.Fatal("Failed to start server")
	}

	log.Println("Server started")

}
