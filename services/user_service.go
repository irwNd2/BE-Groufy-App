package services

import (
	"be-groufy-app/dto/web"
	"be-groufy-app/models"
	"be-groufy-app/repositories"
	"be-groufy-app/utils"
	"errors"
	"time"

	// "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserService struct {
	Repo *repositories.UserRepository
}

// var validate = validator.New()

func (s *UserService) Register(ctx *fiber.Ctx) error {
	user := new(models.User)
	err := ctx.BodyParser(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "internal server error"})
	}

	hashedPass, _ := utils.HashPassword(*user.Password)
	user.Password = &hashedPass

	err = s.Repo.Create(user)
	if err != nil {
		if errors.Is(err, gorm.ErrCheckConstraintViolated) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "fill all the required form"})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "could not create user"})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "registered successfully"})
}

func (s *UserService) GetAllUser(ctx *fiber.Ctx) error {
	users, err := s.Repo.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "internal server error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success get user", "data": users})
}

func (s *UserService) GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := s.Repo.GetById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "ise"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": user})
}

func (s *UserService) GetUserByRole(ctx *fiber.Ctx) error {
	role := ctx.Params("role")
	users, err := s.Repo.GetByRole(role)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "ise"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": users})
}

// best practice
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
