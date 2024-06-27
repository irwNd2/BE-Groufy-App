package repositories

import (
	"be-groufy-app/dto/web"
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

func (r *UserRepository) GetByRole(role string) (users []models.User, err error) {
	err = r.DB.Where("role = ?", role).Find(&users).Error
	return
}

func (r *UserRepository) GetById(id string) (user models.User, err error) {
	err = r.DB.Where("id = ?", id).First(&user).Error
	return
}

func (r *UserRepository) Login(p *web.LoginPayload) (user models.User, err error) {
	err = r.DB.Where("email = ?", p.Email).Find(&user).Error
	return
}
