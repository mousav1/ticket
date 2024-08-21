package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mousv1/ticket/internal/api"
	"github.com/mousv1/ticket/internal/api/handlers"
	"github.com/mousv1/ticket/internal/api/middleware"
)

func SetupRoutes(server *api.Server) error {

	server.App.Get("/", func(c *fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	server.App.Post("/register", handlers.NewUserHandler(server.Store, server.TokenMaker, server.Config).RegisterUser)
	server.App.Post("/login", handlers.NewUserHandler(server.Store, server.TokenMaker, server.Config).LoginUser)

	// Grouped routes that require authentication
	authGroup := server.App.Group("/user", middleware.AuthMiddleware(server.TokenMaker))

	authGroup.Get("/info", func(c *fiber.Ctx) error {
		return c.SendString("test")
	})

	return nil
}
