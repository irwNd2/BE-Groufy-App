package app

import (
	"be-groufy-app/handlers"
	"be-groufy-app/repositories"
	"be-groufy-app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupInfoRoutes(app *fiber.App, db *gorm.DB) {
	handler := &handlers.InfoHandler{
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
