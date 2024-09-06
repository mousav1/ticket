package db

import (
	"context"
	"fmt"
	"time"
)

type reserveSeatResponse struct {
	TicketID   int32     `json:"ticket_id"`
	BusID      int32     `json:"bus_id"`
	SeatID     int32     `json:"seat_id"`
	ReservedAt time.Time `json:"reserved_at"`
}

type ReserveTicketTxParams struct {
	UserID int32 `json:"user_id"`
	BusID  int32 `json:"bus_id"`
	SeatID int32 `json:"seat_id"`
}

func (store *Store) ReserveTicketTx(ctx context.Context, arg ReserveTicketTxParams) (reserveSeatResponse, error) {
	var result reserveSeatResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Reserve the seat by creating a ticket
		ticket, err := q.ReserveTicket(ctx, ReserveTicketParams{
			UserID:            arg.UserID,
			BusID:             arg.BusID,
			SeatReservationID: arg.SeatID,
		})
		if err != nil {
			return fmt.Errorf("failed to reserve ticket: %v", err)
		}

		// Update seat status to reserved
		err = q.UpdateSeatReservationStatus(ctx, UpdateSeatReservationStatusParams{
			BusSeatID: arg.SeatID,
			Status:    "reserved",
			UserID:    arg.UserID,
		})
		if err != nil {
			return fmt.Errorf("failed to update seat status: %v", err)
		}

		// Populate the response
		result = reserveSeatResponse{
			TicketID:   ticket.ID,
			BusID:      ticket.BusID,
			SeatID:     ticket.SeatReservationID,
			ReservedAt: ticket.ReservedAt.Time,
		}

		return nil
	})

	return result, err
}
