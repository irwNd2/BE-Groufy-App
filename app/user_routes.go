package app

import (
	"be-groufy-app/handlers"
	"be-groufy-app/repositories"
	"be-groufy-app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupUserRoutes(app *fiber.App, db *gorm.DB) {
	handler := &handlers.UserHandler{
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
