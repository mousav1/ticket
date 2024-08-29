package routes

import (
	"github.com/mousv1/ticket/internal/api"
	"github.com/mousv1/ticket/internal/api/handlers"
	"github.com/mousv1/ticket/internal/api/middleware"
)

func SetupRoutes(server *api.Server) error {

	server.App.Post("/register", handlers.NewUserHandler(server.Store, server.TokenMaker, server.Config).RegisterUser)
	server.App.Post("/login", handlers.NewUserHandler(server.Store, server.TokenMaker, server.Config).LoginUser)
	server.App.Post("/tokens/renew_access", handlers.NewTokenHandler(server.Store, server.TokenMaker, server.Config).RenewAccessToken)

	// Grouped routes that require authentication
	authGroup := server.App.Group("/user", middleware.AuthMiddleware(server.TokenMaker))
	authGroup.Get("/info", handlers.NewUserHandler(server.Store, server.TokenMaker, server.Config).GetUserProfile)
	authGroup.Put("/update", handlers.NewUserHandler(server.Store, server.TokenMaker, server.Config).UpdateUserProfile)
	authGroup.Post("/password_change", handlers.NewUserHandler(server.Store, server.TokenMaker, server.Config).ChangePassword)

	return nil
}
