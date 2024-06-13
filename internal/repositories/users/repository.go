package users

import (
	"arbitrage-betting/internal/models"
	"fmt"
	"gorm.io/gorm"
)

func NewUsersRepository(db *gorm.DB) Repository {
	return &usersRepository{db: db}
}

type usersRepository struct {
	db *gorm.DB
}

func (r *usersRepository) AddUser(user models.User) (*models.User, error) {
	return &user, nil
	//
	// result := r.db.Create(&user)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	//
	// return &user, nil
}

func (r *usersRepository) GetUserById(id string) (*models.User, error) {
	var user *models.User

	result := r.db.First(&user, "id = ?", fmt.Sprint(id))
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (r *usersRepository) GetUserByName(name string) (*models.User, error) {
	var user *models.User

	result := r.db.First(&user, "name = ?", fmt.Sprint(name))
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (r *usersRepository) UpdateUser(user models.User) error {
	userInDB, err := r.GetUserById(user.ID)
	if err != nil {
		return err
	}

	userInDB.Name = user.Name

	err = r.db.Save(&userInDB).Error
	return err
}

func (r *usersRepository) DeleteUser(id string) error {
	result := r.db.Delete(&models.User{}, "id = ?", id)
	return result.Error
}
