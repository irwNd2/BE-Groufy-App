package handlers

import (
	"be-groufy-app/repositories"
	"be-groufy-app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type InfoHandler struct {
	Service *services.InfoService
}

func (h *InfoHandler) AddInfo(ctx *fiber.Ctx) error {
	return h.Service.AddInfo(ctx)
}

func (h *InfoHandler) GetAllInfo(ctx *fiber.Ctx) error {
	return h.Service.GetAllInfo(ctx)
}

func (h *InfoHandler) GetInfoById(ctx *fiber.Ctx) error {
	return h.Service.GetInfoById(ctx)
}

func (h *InfoHandler) DeleteInfoById(ctx *fiber.Ctx) error {
	return h.Service.DeleteInfoById(ctx)
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
