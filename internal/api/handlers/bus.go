package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	db "github.com/mousv1/ticket/internal/db/sqlc"
	"github.com/mousv1/ticket/internal/token"
	"github.com/mousv1/ticket/internal/util"
)

type BusHandler struct {
	store      *db.Store
	tokenMaker token.Maker
	config     util.Config
}

// Request structure for listing available seats
type listAvailableSeatsRequest struct {
	BusID   int32 `params:"bus_id" validate:"required"`
	RouteID int32 `params:"route_id" validate:"required"`
}

// Response structure for available seats
type seatResponse struct {
	SeatID     int32  `json:"seat_id"`
	SeatNumber int32  `json:"seat_number"`
	Status     string `json:"status"`
}

// Request structure for reserving a seat
type reserveSeatRequest struct {
	RouteID               int32  `json:"route_id" validate:"required"`
	BusID                 int32  `json:"bus_id" validate:"required"`
	SeatID                int32  `json:"seat_id" validate:"required"`
	PassengerNationalCode string `json:"passenger_national_code" validate:"required"`
}

// Response structure for seat reservation
type reserveSeatResponse struct {
	TicketID   int32     `json:"ticket_id"`
	BusID      int32     `json:"bus_id"`
	SeatID     int32     `json:"seat_id"`
	ReservedAt time.Time `json:"reserved_at"`
}

// NewBusHandler creates a new BusHandler
func NewBusHandler(Store *db.Store, tokenMaker token.Maker, Config util.Config) *BusHandler {
	return &BusHandler{
		Store,
		tokenMaker,
		Config,
	}
}

// ListAvailableSeats handles the request to list available seats for a specific bus
func (h *BusHandler) ListAvailableSeats(c *fiber.Ctx) error {
	// Parse query parameters
	var req listAvailableSeatsRequest
	if err := c.ParamsParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request parameters"})
	}

	// Validate request parameters
	validate := validator.New() // Consider using a globally defined validator
	if err := validate.Struct(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Validate that the bus belongs to the specified route
	_, err := h.store.CheckBusRouteAssociation(c.Context(), db.CheckBusRouteAssociationParams{
		ID:   req.RouteID,
		ID_2: req.BusID,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Bus or Route not found or they do not match"})
		}
		// Log error here for debugging if necessary
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to validate bus and route association"})
	}

	// Fetch available seats from the database
	seats, err := h.store.GetAvailableSeatsForBus(c.Context(), db.GetAvailableSeatsForBusParams{
		RouteID: req.RouteID,
		BusID:   req.BusID,
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch available seats"})
	}

	// Map database results to response format
	var response []seatResponse
	for _, seat := range seats {
		response = append(response, seatResponse{
			SeatID:     seat.SeatID,
			SeatNumber: seat.SeatNumber,
			Status:     "available", // Assuming seats returned are available
		})
	}

	// Send response
	return c.Status(http.StatusOK).JSON(response)
}

// ReserveSeat handles the request to reserve a seat on a specific bus
func (h *BusHandler) ReserveSeat(c *fiber.Ctx) error {
	// Parse request body
	var req reserveSeatRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request parameters"})
	}

	// Validate request parameters
	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if the bus is associated with the route
	_, err := h.store.CheckBusRouteAssociation(c.Context(), db.CheckBusRouteAssociationParams{
		ID:   req.RouteID,
		ID_2: req.BusID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(http.StatusNotFound, "Bus or Route not found or they do not match")
		}
		return fiber.NewError(http.StatusInternalServerError, "Failed to validate bus and route association")
	}

	// Check if the seat exists and is available
	seat, err := h.store.GetSeatByID(c.Context(), db.GetSeatByIDParams{
		ID:    req.SeatID,
		BusID: req.BusID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(http.StatusNotFound, "Seat not found")
		}
		return fiber.NewError(http.StatusInternalServerError, "Failed to fetch seat information")
	}

	// Ensure the seat belongs to the bus
	if seat.BusID != req.BusID {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Seat does not belong to the specified bus"})
	}

	// Ensure the seat is available
	if seat.Status != 0 { // Assuming 0 means available
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Seat is not available"})
	}

	payload := c.Locals("authorizationPayloadKey").(*token.Payload)

	user, err := h.store.GetUserByUsername(c.Context(), payload.Username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	// Use the ReserveTicketTx to handle the reservation in a transaction
	reservation, err := h.store.ReserveTicketTx(c.Context(), db.ReserveTicketTxParams{
		UserID: user.ID,
		BusID:  req.BusID,
		SeatID: req.SeatID,
	})
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Failed to reserve seat: "+err.Error())
	}

	// Respond with ticket information
	return c.Status(http.StatusOK).JSON(reservation)
}
