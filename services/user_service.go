package services

import (
	"be-groufy-app/dto/web"
	"be-groufy-app/models"
	"be-groufy-app/repositories"
	"be-groufy-app/utils"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (s *UserService) Register(user *models.User) error {
	hashedPass, _ := utils.HashPassword(*user.Password)
	user.Password = &hashedPass

	err := s.Repo.Create(user)
	return err
}

func (s *UserService) GetAllUser() ([]models.User, error) {
	users, err := s.Repo.GetAll()
	return users, err
}

func (s *UserService) GetUserById(id string) (*models.User, error) {
	user, err := s.Repo.GetById(id)
	return &user, err
}

func (s *UserService) GetUserByRole(role string) ([]models.User, error) {
	users, err := s.Repo.GetByRole(role)
	return users, err
}

func (s *UserService) Login(payload *web.LoginPayload) (res *web.LoginResponse, err error) {
	user, _ := s.Repo.Login(payload)
	if user.ID == 0 {
		return res, errors.New("record_not_found")
	}

	err = utils.ComparePassword([]byte(*user.Password), payload.Password)
	if err != nil {
		return nil, err
	}

	userClaims := web.Claims{
		ID:    user.ID,
		Email: *user.Email,
		Name:  *user.Name,
		Role:  *user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	signedAccessToken, err := utils.NewAccessToken(userClaims)

	if err != nil {
		return res, errors.New("access_token_error")
	}
	result := web.LoginResponse{
		AccessToken: signedAccessToken,
	}
	return &result, err
}
