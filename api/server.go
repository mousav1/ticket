package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	db "github.com/mousv1/ticket/db/sqlc"
	"github.com/mousv1/ticket/token"
	"github.com/mousv1/ticket/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	Config     util.Config
	Store      *db.Queries
	TokenMaker token.Maker
	App        *fiber.App
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store *db.Queries) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TOKENSECRETKEY)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	app := fiber.New()

	server := &Server{
		Config:     config,
		Store:      store,
		TokenMaker: tokenMaker,
		App:        app,
	}

	return server, nil

}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.App.Listen(fmt.Sprintf(":%s", address))
}
