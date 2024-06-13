package users

import (
	"arbitrage-betting/internal/managers/users"
	"github.com/gofiber/fiber/v3"
)

func NewUpdateUserHandler(usersManager users.Manager) *updateUserHandler {
	return &updateUserHandler{usersManager: usersManager}
}

type updateUserHandler struct {
	usersManager users.Manager
}

func (h *updateUserHandler) Serve(c fiber.Ctx) error {
	// var payload models.User

	// data := c.Body()
	//
	// err := json.Unmarshal(data, &payload)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	// }
	//
	// err = h.usersManager.UpdateUser(payload)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	// }

	return c.SendStatus(fiber.StatusOK)
}
