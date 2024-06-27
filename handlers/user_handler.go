package handlers

import (
	"be-groufy-app/dto/web"
	"be-groufy-app/models"
	"be-groufy-app/repositories"
	"be-groufy-app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

	// return h.Service.Register(ctx)
}

func (h *UserHandler) GetAllUser(ctx *fiber.Ctx) error {
	users, err := h.Service.GetAllUser()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "ise"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": users})
	// return h.Service.GetAllUser(ctx)
}

func (h *UserHandler) GetUserByRole(ctx *fiber.Ctx) error {
	role := ctx.Params("role")
	users, err := h.Service.GetUserByRole(role)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "ise"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": users})
	// return h.Service.GetUserByRole(ctx)
}

func (h *UserHandler) GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := h.Service.GetUserById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "ise"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": user})
	// return h.Service.GetUserById(ctx)
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
