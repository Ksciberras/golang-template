package server

import (
	"golang-project-template/internal/config"
	"golang-project-template/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	Router *fiber.App
}

func dbCheckOrigin(origin string) bool {
	return true
}

func Init(cfg config.Config, h *handler.Handler) *Server {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: dbCheckOrigin, AllowCredentials: true,
	}))

	api := app.Group("/api")

	api.Use(logger.New())

	return &Server{
		Router: app,
	}
}
