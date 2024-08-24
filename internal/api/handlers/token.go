package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	db "github.com/mousv1/ticket/internal/db/sqlc"
	"github.com/mousv1/ticket/internal/token"
	"github.com/mousv1/ticket/internal/util"
)

type TokenHandler struct {
	Store      *db.Store
	tokenMaker token.Maker
	Config     util.Config
}

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func NewTokenHandler(Store *db.Store, tokenMaker token.Maker, Config util.Config) *TokenHandler {
	return &TokenHandler{
		Store,
		tokenMaker,
		Config,
	}
}
func (h *TokenHandler) RenewAccessToken(c *fiber.Ctx) error {
	var req renewAccessTokenRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	refreshPayload, err := h.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
	}

	session, err := h.Store.GetSession(c.Context(), refreshPayload.ID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "not found"})
	}

	if session.IsBlocked {
		err := fmt.Errorf("blocked session")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	if session.Username != refreshPayload.Username {
		err := fmt.Errorf("incorrect session user")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("mismatched session token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("expired session")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	accessToken, accessPayload, err := h.tokenMaker.CreateToken(refreshPayload.Username, h.Config.AccessTokenDuration)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create token"})
	}

	rsp := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"response": rsp})
}
