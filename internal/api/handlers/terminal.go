package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	db "github.com/mousv1/ticket/internal/db/sqlc"
	"github.com/mousv1/ticket/internal/token"
	"github.com/mousv1/ticket/internal/util"
)

type TerminalHandler struct {
	store      *db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewTerminalHandler(Store *db.Store, tokenMaker token.Maker, Config util.Config) *TerminalHandler {
	return &TerminalHandler{
		Store,
		tokenMaker,
		Config,
	}
}

func (h *TerminalHandler) ListTerminals(c *fiber.Ctx) error {
	terminals, err := h.store.ListTerminals(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch terminals"})
	}
	return c.JSON(terminals)
}
