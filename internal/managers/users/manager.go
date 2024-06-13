package users

import (
	"arbitrage-betting/internal/models"
	"arbitrage-betting/internal/repositories/users"
)

func NewUserManager(usersRepository users.Repository) Manager {
	return &userManager{usersRepository: usersRepository}
}

type userManager struct {
	usersRepository users.Repository
}

func (m *userManager) AddUser(user models.User) (*models.User, error) {
	return m.usersRepository.AddUser(user)
}

func (m *userManager) GetUserById(s string) (*models.User, error) {
	return m.usersRepository.GetUserById(s)
}

func (m *userManager) GetUserByName(s string) (*models.User, error) {
	return m.usersRepository.GetUserByName(s)
}

func (m *userManager) UpdateUser(user models.User) error {
	return m.usersRepository.UpdateUser(user)
}

func (m *userManager) DeleteUser(id string) error {
	return m.usersRepository.DeleteUser(id)
}
