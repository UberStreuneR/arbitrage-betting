package users

import (
	"arbitrage-betting/internal/managers/users"
	"github.com/gofiber/fiber/v3"
)

func NewGetUserByNameHandler(usersManager users.Manager) *getUserByNameHandler {
	return &getUserByNameHandler{usersManager: usersManager}
}

type getUserByNameHandler struct {
	usersManager users.Manager
}

func (h *getUserByNameHandler) Serve(c fiber.Ctx) error {
	name := c.Query("name")

	if name == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}

	user, err := h.usersManager.GetUserByName(name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if user.Name == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": fiber.Map{"user": user}})
}
