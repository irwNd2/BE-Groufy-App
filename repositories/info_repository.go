package repositories

import (
	"be-groufy-app/models"

	"gorm.io/gorm"
)

type InfoRepository struct {
	DB *gorm.DB
}

func (r *InfoRepository) Create(info *models.Info) (*models.Info, error) {
	err := r.DB.Create(info).Error
	return info, err
}

func (r *InfoRepository) GetAll() ([]models.Info, error) {
	var infos []models.Info
	err := r.DB.Find(&infos).Error
	return infos, err
}

func (r *InfoRepository) GetById(id string) (*models.Info, error) {
	var info models.Info
	err := r.DB.Where("id = ?", id).First(&info).Error
	return &info, err
}

func (r *InfoRepository) DeleteById(id string) error {
	return r.DB.Delete(&models.Info{}, id).Error
}
