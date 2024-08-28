package main

import (
	"context"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/mousv1/ticket/internal/api"
	db "github.com/mousv1/ticket/internal/db/sqlc"
	"github.com/mousv1/ticket/internal/routes"
	"github.com/mousv1/ticket/internal/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.APPDEBUG == "true" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Connect to the database
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUSERNAME,
		config.DBPASSWORD,
		config.DBHOST,
		config.DBPORT,
		config.DBDATABASE,
	)

	connPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	defer connPool.Close()

	runDBMigration(config.MigrationURL, dsn)

	store := db.NewStore(connPool)

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

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}
