package repositories

import (
	"coffee-pos-backend/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), 14)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)
	return r.db.Create(&user).Error
}

func (r *UserRepository) GetUserByID(id uint) (*models.UserWithRole, error) {

	var result models.UserWithRole
	err := r.db.Model(&models.User{}).Select("users.*, roles.role_name").Joins("left join roles on roles.id = users.role_id").Where("users.id = ?", id).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.db.Save(&user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *UserRepository) GetAllUsers() ([]models.UserWithRole, error) {
	var users []models.UserWithRole
	err := r.db.Model(&models.User{}).Select("users.*, roles.role_name").Joins("left join roles on roles.id = users.role_id").Scan(&users).Error
	return users, err
}
