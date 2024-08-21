// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	// buses.sql
	CreateBus(ctx context.Context, arg CreateBusParams) (Bus, error)
	CreateBusSeat(ctx context.Context, arg CreateBusSeatParams) (BusSeat, error)
	CreateCity(ctx context.Context, name string) (City, error)
	CreatePenalty(ctx context.Context, arg CreatePenaltyParams) (Penalty, error)
	// routes.sql
	CreateRoute(ctx context.Context, arg CreateRouteParams) (Route, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	// terminals.sql
	CreateTerminal(ctx context.Context, arg CreateTerminalParams) (Terminal, error)
	// users.sql
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteTicket(ctx context.Context, id int32) error
	GetAllCities(ctx context.Context) ([]City, error)
	// پیدا کردن تمامی مسیرها
	GetAllRoutes(ctx context.Context) ([]Route, error)
	GetBusByID(ctx context.Context, id int32) (Bus, error)
	GetBusPenalties(ctx context.Context, busID pgtype.Int4) ([]Penalty, error)
	GetBusSeats(ctx context.Context, busID pgtype.Int4) ([]BusSeat, error)
	// cities.sql
	GetCityByID(ctx context.Context, id int32) (City, error)
	GetReservedTicketsCount(ctx context.Context, busID pgtype.Int4) (int64, error)
	// پیدا کردن یک مسیر بر اساس مبدا و مقصد
	GetRouteByTerminals(ctx context.Context, arg GetRouteByTerminalsParams) (Route, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetTerminalByID(ctx context.Context, id int32) (Terminal, error)
	GetTerminalsByCity(ctx context.Context, cityID pgtype.Int4) ([]Terminal, error)
	GetTicketByID(ctx context.Context, id int32) (GetTicketByIDRow, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserByID(ctx context.Context, id int32) (GetUserByIDRow, error)
	GetUserTickets(ctx context.Context, userID pgtype.Int4) ([]GetUserTicketsRow, error)
	ReserveTicket(ctx context.Context, arg ReserveTicketParams) error
	SearchBuses(ctx context.Context, arg SearchBusesParams) ([]Bus, error)
	SearchBusesByCities(ctx context.Context, arg SearchBusesByCitiesParams) ([]Bus, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)