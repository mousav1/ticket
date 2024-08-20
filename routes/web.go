package routes

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/mousv1/ticket/db/sqlc"
)

func SetupRoutes(app *fiber.App, store *db.Queries) error {
	// Define a route for the GET method on the root path '/'

	app.Get("/", func(c *fiber.Ctx) error { // Change `c fiber.Ctx` to `c *fiber.Ctx`
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	return nil
}
