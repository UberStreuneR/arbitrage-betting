package users

import (
	"arbitrage-betting/internal/managers/users"
	"github.com/gofiber/fiber/v3"
)

func NewDeleteUserHandler(usersManager users.Manager) *deleteUserHandler {
	return &deleteUserHandler{usersManager: usersManager}
}

type deleteUserHandler struct {
	usersManager users.Manager
}

func (h *deleteUserHandler) Serve(c fiber.Ctx) error {
	// userId := c.Params("id")
	//
	// err := h.usersManager.DeleteUser(userId)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	// }

	return c.SendStatus(fiber.StatusNoContent)
}
