package repositories

import (
	"be-groufy-app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetByRole(role string) ([]models.User, error) {
	var users []models.User
	err := r.DB.Where("role = ?", role).Error
	return users, err
}

func (r *UserRepository) GetById(id string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}
