package repository

import (
	"gogo/internal/app/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(user *entities.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByUsername(username string) (*entities.User, error) {
	var user entities.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
