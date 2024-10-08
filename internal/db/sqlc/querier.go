// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CheckBusRouteAssociation(ctx context.Context, arg CheckBusRouteAssociationParams) (CheckBusRouteAssociationRow, error)
	CheckSeatAvailability(ctx context.Context, arg CheckSeatAvailabilityParams) (CheckSeatAvailabilityRow, error)
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
	GetAllRoutes(ctx context.Context) ([]Route, error)
	GetAvailableSeatsForBus(ctx context.Context, arg GetAvailableSeatsForBusParams) ([]GetAvailableSeatsForBusRow, error)
	GetBusByID(ctx context.Context, id int32) (Bus, error)
	GetBusPenalties(ctx context.Context, busID int32) ([]Penalty, error)
	GetBusSeats(ctx context.Context, busID int32) ([]BusSeat, error)
	// cities.sql
	GetCityByID(ctx context.Context, id int32) (City, error)
	GetReservedTicketsCount(ctx context.Context, busID int32) (int64, error)
	GetRouteByID(ctx context.Context, id int32) (Route, error)
	GetSeatByID(ctx context.Context, arg GetSeatByIDParams) (GetSeatByIDRow, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetTerminalByID(ctx context.Context, id int32) (Terminal, error)
	GetTerminalsByCity(ctx context.Context, cityID int32) ([]Terminal, error)
	GetTicketByID(ctx context.Context, id int32) (GetTicketByIDRow, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserByID(ctx context.Context, id int32) (GetUserByIDRow, error)
	GetUserByUsername(ctx context.Context, username string) (GetUserByUsernameRow, error)
	GetUserTickets(ctx context.Context, userID int32) ([]GetUserTicketsRow, error)
	ListRoutes(ctx context.Context, arg ListRoutesParams) ([]ListRoutesRow, error)
	ListTerminals(ctx context.Context) ([]ListTerminalsRow, error)
	ListUserTickets(ctx context.Context, userID int32) ([]ListUserTicketsRow, error)
	PurchaseTicket(ctx context.Context, arg PurchaseTicketParams) (PurchaseTicketRow, error)
	ReserveTicket(ctx context.Context, arg ReserveTicketParams) (ReserveTicketRow, error)
	SearchBuses(ctx context.Context, arg SearchBusesParams) ([]Bus, error)
	SearchBusesByCities(ctx context.Context, arg SearchBusesByCitiesParams) ([]Bus, error)
	UpdateSeatReservationStatus(ctx context.Context, arg UpdateSeatReservationStatusParams) error
	// Ensures no conflicting reservation or purchase exists
	UpdateSeatStatusAfterTrip(ctx context.Context, busID int32) error
	UpdateTicketStatus(ctx context.Context, arg UpdateTicketStatusParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error)
}

var _ Querier = (*Queries)(nil)
