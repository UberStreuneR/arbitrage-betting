package main

import (
	"arbitrage-betting/internal/repositories/users"
	"gorm.io/gorm"
)

type repositories struct {
	users users.Repository
}

func setupRepositories() *repositories {
	// TODO:: gorm db

	var db *gorm.DB

	return &repositories{
		users: users.NewUsersRepository(db),
	}
}
