package repositories

import (
	"coffee-pos-backend/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
}

type userRepository struct {
 db *gorm.DB
}


func NewUserRepository (db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}