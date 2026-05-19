package repositories

import (
	"backend/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	UpdateRefreshToken(id string, refreshToken string) error
	UpdateUser(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) UpdateRefreshToken(id string, refreshToken string) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("refresh_token", refreshToken).Error
}

func (r *userRepository) UpdateUser(user *models.User) error {
	if err := r.db.Save(&user).Error; err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	return nil
}
