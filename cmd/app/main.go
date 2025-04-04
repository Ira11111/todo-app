package main

import (
	"github.com/Ira11111/todo-app/internal/config"
	"github.com/Ira11111/todo-app/internal/handlers"
	"github.com/Ira11111/todo-app/internal/repository"
	"github.com/Ira11111/todo-app/internal/services"
	"github.com/Ira11111/todo-app/internal/transport"
	"log"
)

func main() {
	cfg := config.MustLoad()

	appDB, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	appRepos := repository.NewRepository(appDB)
	appServices := services.NewService(appRepos)
	appHandlers := handlers.NewHandler(appServices)
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
