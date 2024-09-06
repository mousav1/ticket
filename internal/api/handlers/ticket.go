package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	db "github.com/mousv1/ticket/internal/db/sqlc"
	"github.com/mousv1/ticket/internal/token"
	"github.com/mousv1/ticket/internal/util"
)

type TicketHandler struct {
	store      *db.Store
	tokenMaker token.Maker
	config     util.Config
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

// ListUserTicketsResponse represents the response structure for user tickets
type ListUserTicketsResponse struct {
	TicketID      int32     `json:"ticket_id"`
	BusID         int32     `json:"bus_id"`
	SeatID        int32     `json:"seat_id"`
	ReservedAt    time.Time `json:"reserved_at"`
	DepartureTime time.Time `json:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time"`
	Price         int       `json:"price"`
	SeatNumber    int       `json:"seat_number"`
	Status        string    `json:"status"`
}

// NewTicketHandler creates a new TicketHandler
func NewTicketHandler(Store *db.Store, tokenMaker token.Maker, Config util.Config) *TicketHandler {
	return &TicketHandler{
		Store,
		tokenMaker,
		Config,
	}
}

// ReserveSeat handles the request to reserve a seat on a specific bus
func (h *TicketHandler) ReserveSeat(c *fiber.Ctx) error {
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
	seat, err := h.store.CheckSeatAvailability(c.Context(), db.CheckSeatAvailabilityParams{
		ID:    req.SeatID,
		BusID: req.BusID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(http.StatusNotFound, "Seat not found")
		}
		return fiber.NewError(http.StatusInternalServerError, "Failed to fetch seat information")
	}

	// Ensure the seat is available for purchase
	if seat.Status != "available" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Seat is not available for purchase"})
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

// PurchaseTicket handles the request to purchase a ticket
func (h *TicketHandler) PurchaseTicket(c *fiber.Ctx) error {
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
	seat, err := h.store.CheckSeatAvailability(c.Context(), db.CheckSeatAvailabilityParams{
		ID:    req.SeatID,
		BusID: req.BusID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(http.StatusNotFound, "Seat not found")
		}
		return fiber.NewError(http.StatusInternalServerError, "Failed to fetch seat information")
	}

	// Ensure the seat is available for purchase
	if seat.Status != "available" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Seat is not available for purchase"})
	}

	// Get user from context
	payload := c.Locals("authorizationPayloadKey").(*token.Payload)
	user, err := h.store.GetUserByUsername(c.Context(), payload.Username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	// Start the transaction to purchase the ticket
	result, err := h.store.PurchaseTicketTx(c.Context(), db.PurchaseTicketTxParams{
		UserID: user.ID,
		BusID:  req.BusID,
		SeatID: req.SeatID,
	})
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Failed to purchase ticket")
	}

	// Respond with ticket information
	response := reserveSeatResponse{
		TicketID:   result.TicketID,
		BusID:      result.BusID,
		SeatID:     result.SeatID,
		ReservedAt: result.ReservedAt,
	}

	return c.Status(http.StatusOK).JSON(response)
}

// ListUserTickets handles the request to list all tickets of a user
func (h *TicketHandler) ListUserTickets(c *fiber.Ctx) error {
	// Extract user info from authorization payload
	payload := c.Locals("authorizationPayloadKey").(*token.Payload)

	// Fetch user by username
	user, err := h.store.GetUserByUsername(c.Context(), payload.Username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Fetch tickets for the user
	tickets, err := h.store.ListUserTickets(c.Context(), user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user tickets"})
	}

	// Map tickets to response format
	var response []ListUserTicketsResponse
	for _, ticket := range tickets {
		response = append(response, ListUserTicketsResponse{
			TicketID:      ticket.TicketID,
			BusID:         ticket.BusID,
			SeatID:        ticket.SeatID,
			ReservedAt:    ticket.ReservedAt.Time,
			DepartureTime: ticket.DepartureTime,
			ArrivalTime:   ticket.ArrivalTime,
			Price:         int(ticket.Price),
			SeatNumber:    int(ticket.SeatNumber),
			Status:        ticket.ReservationStatus,
		})
	}

	// Send the response
	return c.Status(http.StatusOK).JSON(response)
}

// CancelTicket handles the request to cancel a ticket
func (h *TicketHandler) CancelTicket(c *fiber.Ctx) error {
	// Parse ticket ID from URL parameters
	ticketID, err := strconv.Atoi(c.Params("id"))
	if err != nil || ticketID <= 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ticket ID"})
	}

	// Extract user information from token payload
	payload := c.Locals("authorizationPayloadKey").(*token.Payload)
	user, err := h.store.GetUserByUsername(c.Context(), payload.Username)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Retrieve the ticket details
	ticket, err := h.store.GetTicketByID(c.Context(), int32(ticketID))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Ticket not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch ticket details"})
	}

	// Check if the ticket belongs to the current user
	if ticket.UserID != user.ID {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "You do not have permission to cancel this ticket"})
	}

	// Check if the ticket is already canceled or non-cancelable
	if ticket.Status != "canceled" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Ticket cannot be canceled"})
	}

	// Start a transaction to cancel the ticket
	err = h.store.CancelTicketTx(c.Context(), db.CancelTicketParams{
		TicketID: ticket.ID,
		SeatID:   ticket.SeatReservationID,
		UserID:   user.ID,
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to cancel ticket"})
	}

	// Respond with success message
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Ticket canceled successfully"})
}
