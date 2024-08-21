package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	db "github.com/mousv1/ticket/internal/db/sqlc"
	"github.com/mousv1/ticket/internal/token"
	"github.com/mousv1/ticket/internal/util"
)

type UserHandler struct {
	Store      *db.Queries
	tokenMaker token.Maker
	Config     util.Config
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
}

type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type loginUserResponse struct {
	SessionID             uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  userResponse `json:"user"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func NewUserHandler(Store *db.Queries, tokenMaker token.Maker, Config util.Config) *UserHandler {
	return &UserHandler{
		Store,
		tokenMaker,
		Config,
	}
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not hash password"})
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
	}

	user, err := h.Store.CreateUser(c.Context(), arg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create user"})
	}

	if err != nil {
		// Assuming db.ErrorCode(err) is used to get the specific error code
		if db.ErrorCode(err) == db.UniqueViolation {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "username already exists"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create user"})
	}

	rsp := newUserResponse(user)
	return c.Status(fiber.StatusCreated).JSON(rsp)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

func (h *UserHandler) LoginUser(c *fiber.Ctx) error {
	var req loginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.Store.GetUser(c.Context(), req.Username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid username or password"})
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid username or password"})
	}

	accessToken, accessPayload, err := h.tokenMaker.CreateToken(user.Username, h.Config.AccessTokenDuration)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create token"})
	}

	refreshToken, refreshPayload, err := h.tokenMaker.CreateToken(user.Username, h.Config.RefreshTokenDuration)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create token"})
	}

	session, err := h.Store.CreateSession(c.Context(), db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    c.Get("User-Agent"),
		ClientIp:     c.IP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "could not create session"})
	}

	rsp := loginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"response": rsp})
}

// GetUserProfile handles fetching user profile
// func (h *UserHandler) GetUserProfile(c *fiber.Ctx) error {
// 	userID := c.Locals("user_id").(int64)

// 	user, err := h.Store.GetUserByID(c.Context(), userID)
// 	if err != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
// 	}

// 	return c.JSON(user)
// }
