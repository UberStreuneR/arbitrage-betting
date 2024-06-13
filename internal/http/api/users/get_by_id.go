package users

import (
	"arbitrage-betting/internal/managers/users"
	"github.com/gofiber/fiber/v3"
)

func NewGetUserByIdHandler(usersManager users.Manager) *getUserByIdHandler {
	return &getUserByIdHandler{usersManager: usersManager}
}

type getUserByIdHandler struct {
	usersManager users.Manager
}

func (h *getUserByIdHandler) Serve(c fiber.Ctx) error {
	// id := c.Params("id")

	return c.SendStatus(fiber.StatusNotFound)

	// user := models.User{"1", "New Username"}
	// user, err := h.usersManager.GetUserById(id)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// if user.Name == "" {
	// 	return c.SendStatus(fiber.StatusNotFound)
	// }
	//
	// return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": fiber.Map{"user": user}})
}
