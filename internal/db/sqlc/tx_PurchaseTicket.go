package db

import (
	"context"
	"fmt"
	"time"
)

type PurchaseTicketTxParams struct {
	UserID int32 `json:"user_id"`
	BusID  int32 `json:"bus_id"`
	SeatID int32 `json:"seat_id"`
}

type PurchaseTicketTxResult struct {
	TicketID   int32     `json:"ticket_id"`
	BusID      int32     `json:"bus_id"`
	SeatID     int32     `json:"seat_id"`
	ReservedAt time.Time `json:"reserved_at"`
}

// PurchaseTicketTx handles purchasing a ticket in a transaction
func (store *Store) PurchaseTicketTx(ctx context.Context, arg PurchaseTicketTxParams) (PurchaseTicketTxResult, error) {
	var result PurchaseTicketTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Reserve the ticket by updating the status to purchased
		ticket, err := q.PurchaseTicket(ctx, PurchaseTicketParams{
			UserID:            arg.UserID,
			BusID:             arg.BusID,
			SeatReservationID: arg.SeatID,
		})
		if err != nil {
			return err
		}

		// Update seat status to purchased
		err = q.UpdateSeatReservationStatus(ctx, UpdateSeatReservationStatusParams{
			BusSeatID: arg.SeatID,
			Status:    "purchased",
			UserID:    arg.UserID,
		})
		if err != nil {
			return err
		}

		result = PurchaseTicketTxResult{
			TicketID:   ticket.ID,
			BusID:      arg.BusID,
			SeatID:     arg.SeatID,
			ReservedAt: ticket.PurchasedAt.Time,
		}

		return nil
	})

	if err != nil {
		return result, fmt.Errorf("failed to purchase ticket: %v", err)
	}

	return result, nil
}
