// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Bus struct {
	ID               int32
	RouteID          pgtype.Int4
	DepartureTime    pgtype.Timestamptz
	ArrivalTime      pgtype.Timestamptz
	Capacity         int32
	Price            int32
	BusType          string
	Corporation      pgtype.Text
	SuperCorporation pgtype.Text
	ServiceNumber    pgtype.Text
	IsVip            pgtype.Bool
}

type BusSeat struct {
	ID                    int32
	BusID                 pgtype.Int4
	SeatNumber            int32
	Status                int32
	PassengerNationalCode pgtype.Text
}

type City struct {
	ID   int32
	Name string
}

type Penalty struct {
	ID                int32
	BusID             pgtype.Int4
	ActualHoursBefore pgtype.Float8
	HoursBefore       pgtype.Float8
	Percent           int32
	CustomText        pgtype.Text
}

type Route struct {
	ID                    int32
	OriginTerminalID      pgtype.Int4
	DestinationTerminalID pgtype.Int4
	Duration              pgtype.Interval
	Distance              int32
}

type Terminal struct {
	ID     int32
	CityID pgtype.Int4
	Name   string
}

type Ticket struct {
	ID         int32
	UserID     pgtype.Int4
	BusID      pgtype.Int4
	SeatID     pgtype.Int4
	ReservedAt pgtype.Timestamptz
}

type User struct {
	ID           int32
	Username     string
	PasswordHash string
}
