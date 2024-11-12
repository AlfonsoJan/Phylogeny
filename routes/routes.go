package routes

import (
	"Phylogeny/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/job", handlers.CreateJobHandler)
}
