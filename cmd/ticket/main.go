package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mousv1/ticket/internal/api"
	db "github.com/mousv1/ticket/internal/db/sqlc"
	"github.com/mousv1/ticket/internal/routes"
	"github.com/mousv1/ticket/internal/util"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	// Connect to the database
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUSERNAME,
		config.DBPASSWORD,
		config.DBHOST,
		config.DBPORT,
		config.DBDATABASE,
	)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to connect to database")
	}
	defer conn.Close(context.Background())

	store := db.New(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create serve")
	}

	// Register the routes
	if err := routes.SetupRoutes(server); err != nil {
		log.Fatal().Err(err).Msg("ailed to set up routes")
	}

	err = server.Start(config.APPPORT)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start serve")
	}
}
