package handlers

import (
	"be-groufy-app/dto/web"
	"be-groufy-app/models"
	"be-groufy-app/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service *services.UserService
}

func (h *UserHandler) RegisterUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	err := ctx.BodyParser(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Bad request"})
	}

	err = h.Service.Register(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "ise"})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "registered successfully"})
}

func (h *UserHandler) GetAllUser(ctx *fiber.Ctx) error {
	users, err := h.Service.GetAllUser()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "ise"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": users})
}

func (h *UserHandler) GetUserByRole(ctx *fiber.Ctx) error {
	role := ctx.Params("role")
	users, err := h.Service.GetUserByRole(role)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "ise"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": users})
}

func (h *UserHandler) GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := h.Service.GetUserById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "ise"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": user})
}

func (h *UserHandler) AuthLogin(ctx *fiber.Ctx) error {
	var input web.LoginPayload
	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "ise"})
	}
	res, err := h.Service.Login(&input)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad request"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": res})
}
