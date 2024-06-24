package handlers

import (
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
}
