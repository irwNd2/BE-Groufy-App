package handlers

import (
	"be-groufy-app/models"
	"be-groufy-app/repositories"
	"be-groufy-app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type InfoHandler struct {
	Service *services.InfoService
}

func (h *InfoHandler) AddInfo(ctx *fiber.Ctx) error {
	info := new(models.Info)
	err := ctx.BodyParser(info)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": "bad request"})
	}
	info, err = h.Service.AddInfo(info)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": "bad request"})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "success create info", "data": info})

}

func (h *InfoHandler) GetAllInfo(ctx *fiber.Ctx) error {
	infos, err := h.Service.GetAllInfo()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "ise"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": infos})
}

func (h *InfoHandler) GetInfoById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	info, err := h.Service.GetInfoById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "ise"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": info})
}

func (h *InfoHandler) DeleteInfoById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := h.Service.DeleteInfoById(id)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "ise"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success delete info"})
}

func SetupInfoRoutes(app *fiber.App, db *gorm.DB) {
	handler := &InfoHandler{
		Service: &services.InfoService{
			Repo: &repositories.InfoRepository{
				DB: db,
			},
		},
	}

	api := app.Group("/api/info")
	api.Post("/add", handler.AddInfo)
	api.Get("/all", handler.GetAllInfo)
	api.Get("/:id", handler.GetInfoById)
	api.Delete("/:id", handler.DeleteInfoById)
}
