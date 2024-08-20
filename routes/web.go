package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mousv1/ticket/api"
	"github.com/mousv1/ticket/api/middleware"
)

func SetupRoutes(server *api.Server) error {

	server.App.Get("/", func(c *fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	// Grouped routes that require authentication
	authGroup := server.App.Group("/api/v1", middleware.AuthMiddleware(server.TokenMaker))

	authGroup.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("test")
	})

	return nil
}
