package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	db "github.com/mousv1/ticket/internal/db/sqlc"
	"github.com/mousv1/ticket/internal/token"
	"github.com/mousv1/ticket/internal/util"
)

type CityHandler struct {
	store      *db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewCityHandler(Store *db.Store, tokenMaker token.Maker, Config util.Config) *CityHandler {
	return &CityHandler{
		Store,
		tokenMaker,
		Config,
	}
}

func (h *CityHandler) ListCities(c *fiber.Ctx) error {
	cities, err := h.store.GetAllCities(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch cities",
		})
	}

	return c.Status(http.StatusOK).JSON(cities)
}
