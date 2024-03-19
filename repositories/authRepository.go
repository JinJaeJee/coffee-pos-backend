package repositories

import (
	"coffee-pos-backend/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}

}

func (repo *AuthRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	result := repo.db.Where("username = ?", username).First(&user)
	return &user, result.Error
}
