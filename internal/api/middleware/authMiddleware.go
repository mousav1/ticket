package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mousav1/ticket/internal/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware creates a gin middleware for authorization
func AuthMiddleware(tokenMaker token.Maker) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract the Authorization header
		authHeader := c.Get(authorizationHeaderKey)
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing authorization header"})
		}

		// Assume Bearer token format: "Bearer <token>"
		tokenStr := authHeader[len("Bearer "):]

		// Verify the token
		payload, err := tokenMaker.VerifyToken(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		// Attach user information to the request context
		c.Locals("authorizationPayloadKey", payload)

		// Proceed to the next handler
		return c.Next()
	}
}
