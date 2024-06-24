package services

import (
	"be-groufy-app/models"
	"be-groufy-app/repositories"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type InfoService struct {
	Repo *repositories.InfoRepository
}

func (s *InfoService) AddInfo(ctx *fiber.Ctx) error {
	info := new(models.Info)
	if err := ctx.BodyParser(info); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "request failed"})
	}
	if err := s.Repo.Create(info); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "could not add information"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "info has been added"})
}

func (s *InfoService) GetAllInfo(ctx *fiber.Ctx) error {
	infos, err := s.Repo.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "could not get info"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "info fetched successfully",
		"data":    infos,
	})
}

func (s *InfoService) GetInfoById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	info, err := s.Repo.GetById(id)
	fmt.Println("err:", err)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Info not found"})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "could not get info"})
	}
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "id cannot be empty"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success get info", "data": info})
}

func (s *InfoService) DeleteInfoById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := s.Repo.DeleteById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "could not delete the info"})
	}
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "id cannot be empty"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "info deleted successfully"})

}
