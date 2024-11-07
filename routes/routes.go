package routes

import (
	"Phylogeny/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/job/:id", handlers.GetJobHandler)
	api.Post("/job", handlers.CreateJobHandler)
	api.Put("/job/:id", handlers.UpdateJobHandler)
	api.Delete("/job/:id", handlers.DeleteJobHandler)
}
