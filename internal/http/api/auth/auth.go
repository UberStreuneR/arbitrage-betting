package auth

import (
	"arbitrage-betting/internal/managers/auth"
	"arbitrage-betting/internal/models"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func NewAuthHandler(config models.AuthConfig, authManager auth.Manager) *authHandler {
	return &authHandler{config: config, authManager: authManager}
}

type authHandler struct {
	config      models.AuthConfig
	authManager auth.Manager
}

func (h *authHandler) Serve(c fiber.Ctx) error {
	data := c.Body()

	var payload models.UserLoginSchema

	err := json.Unmarshal(data, &payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	valid, err := h.authManager.Authenticate(payload.Username, payload.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if !valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims := jwt.MapClaims{
		"name":  payload.Username,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
