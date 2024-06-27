package services

import (
	"be-groufy-app/models"
	"be-groufy-app/repositories"
)

type InfoService struct {
	Repo *repositories.InfoRepository
}

func (s *InfoService) AddInfo(info *models.Info) (*models.Info, error) {
	info, err := s.Repo.Create(info)
	return info, err
}

func (s *InfoService) GetAllInfo() ([]models.Info, error) {
	infos, err := s.Repo.GetAll()
	return infos, err
}

func (s *InfoService) GetInfoById(id string) (*models.Info, error) {
	info, err := s.Repo.GetById(id)
	return info, err
}

func (s *InfoService) DeleteInfoById(id string) error {
	return s.Repo.DeleteById(id)
}
