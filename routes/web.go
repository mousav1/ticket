package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mousv1/ticket/api"
)

func SetupRoutes(server *api.Server) error {

	server.App.Get("/", func(c *fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	return nil
}
