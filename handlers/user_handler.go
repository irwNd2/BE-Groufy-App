package handlers

import (
	"be-groufy-app/dto/web"
	"be-groufy-app/repositories"
	"be-groufy-app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct {
	Service *services.UserService
}

func (h *UserHandler) RegisterUser(ctx *fiber.Ctx) error {
	return h.Service.Register(ctx)
}

func (h *UserHandler) GetAllUser(ctx *fiber.Ctx) error {
	return h.Service.GetAllUser(ctx)
}

func (h *UserHandler) GetUserByRole(ctx *fiber.Ctx) error {
	return h.Service.GetUserByRole(ctx)
}

func (h *UserHandler) GetUserById(ctx *fiber.Ctx) error {
	return h.Service.GetUserById(ctx)
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

func SetupUserRoutes(app *fiber.App, db *gorm.DB) {
	handler := &UserHandler{
		Service: &services.UserService{
			Repo: &repositories.UserRepository{
				DB: db,
			},
		},
	}

	api := app.Group("/api/user")
	api.Post("/register", handler.RegisterUser)
	api.Get("/all", handler.GetAllUser)
	api.Get("/:id", handler.GetUserById)
	api.Get("/role/:role", handler.GetUserByRole)
	api.Post("/login", handler.AuthLogin)
}
