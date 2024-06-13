package users

import "arbitrage-betting/internal/models"

type Repository interface {
	AddUser(models.User) (*models.User, error)
	GetUserById(string) (*models.User, error)
	GetUserByName(string) (*models.User, error)
	UpdateUser(models.User) error
	DeleteUser(string) error
}
