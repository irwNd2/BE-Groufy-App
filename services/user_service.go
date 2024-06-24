package services

import (
	"be-groufy-app/models"
	"be-groufy-app/repositories"

	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (s *UserService) Register(ctx *fiber.Ctx) error {
	user := new(models.User)
	err := ctx.BodyParser(user)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "internal server error"})
	}
	err = s.Repo.Create(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "could not create user"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "registered successfully"})
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
	user, err := s.Repo.GetByRole(role)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "ise"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": user})
}


