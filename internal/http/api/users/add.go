package users

import (
	"arbitrage-betting/internal/managers/users"
	"arbitrage-betting/internal/models"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"strings"
)

func NewAddUserHandler(usersManager users.Manager) *addUserHandler {
	return &addUserHandler{usersManager: usersManager}
}

type addUserHandler struct {
	usersManager users.Manager
}

func (h *addUserHandler) Serve(c fiber.Ctx) error {
	var payload models.User

	data := c.Body()

	err := json.Unmarshal(data, &payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.usersManager.AddUser(payload)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User with such name already exists"})
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": fiber.Map{"user": user}})
}
