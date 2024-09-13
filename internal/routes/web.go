package routes

import (
	"github.com/mousav1/ticket/internal/api"
	"github.com/mousav1/ticket/internal/api/handlers"
	"github.com/mousav1/ticket/internal/api/middleware"
)

func SetupRoutes(server *api.Server) error {

	server.App.Post("/register", handlers.NewUserHandler(server.Store, server.TokenMaker, server.Config).RegisterUser)
	server.App.Post("/login", handlers.NewUserHandler(server.Store, server.TokenMaker, server.Config).LoginUser)
	server.App.Post("/tokens/renew_access", handlers.NewTokenHandler(server.Store, server.TokenMaker, server.Config).RenewAccessToken)
	server.App.Get("/cities", handlers.NewCityHandler(server.Store, server.TokenMaker, server.Config).ListCities)
	server.App.Get("/terminals", handlers.NewTerminalHandler(server.Store, server.TokenMaker, server.Config).ListTerminals)
	server.App.Get("/routes", handlers.NewRouteHandler(server.Store, server.TokenMaker, server.Config).SearchRoutes)
	server.App.Get("/routes/:route_id/buses/:bus_id/seats", handlers.NewBusHandler(server.Store, server.TokenMaker, server.Config).ListAvailableSeats)
	// Grouped routes that require authentication
	authGroup := server.App.Group("/", middleware.AuthMiddleware(server.TokenMaker))
	authGroup.Get("/user/info", handlers.NewUserHandler(server.Store, server.TokenMaker, server.Config).GetUserProfile)
	authGroup.Put("/user/update", handlers.NewUserHandler(server.Store, server.TokenMaker, server.Config).UpdateUserProfile)
	authGroup.Post("/user/password_change", handlers.NewUserHandler(server.Store, server.TokenMaker, server.Config).ChangePassword)
	authGroup.Get("/user/tickets", handlers.NewTicketHandler(server.Store, server.TokenMaker, server.Config).ListUserTickets)
	authGroup.Get("/routes/reserve", handlers.NewTicketHandler(server.Store, server.TokenMaker, server.Config).ReserveSeat)
	authGroup.Get("/routes/purchase", handlers.NewTicketHandler(server.Store, server.TokenMaker, server.Config).PurchaseTicket)
	authGroup.Delete("/tickets/:id", handlers.NewTicketHandler(server.Store, server.TokenMaker, server.Config).CancelTicket)
	return nil
}
