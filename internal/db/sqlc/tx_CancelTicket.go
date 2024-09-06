package db

import (
	"context"
	"fmt"
)

// CancelTicketParams holds the input parameters for CancelTicketTx
type CancelTicketParams struct {
	UserID   int32 `json:"user_id"`
	TicketID int32 `json:"ticket_id"`
	SeatID   int32 `json:"seat_id"`
}

// CancelTicketTx cancels a ticket and updates the seat status within a transaction
func (store *Store) CancelTicketTx(ctx context.Context, arg CancelTicketParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		// Update the ticket status to canceled
		err := q.UpdateTicketStatus(ctx, UpdateTicketStatusParams{
			ID:     arg.TicketID,
			Status: "canceled",
		})
		if err != nil {
			return fmt.Errorf("failed to update ticket status: %w", err)
		}

		// Update the seat status to available
		err = q.UpdateSeatReservationStatus(ctx, UpdateSeatReservationStatusParams{
			BusSeatID: arg.SeatID,
			Status:    "canceled",
			UserID:    arg.UserID,
		})
		if err != nil {
			return fmt.Errorf("failed to update seat status: %w", err)
		}

		return nil
	})

	return err
}
