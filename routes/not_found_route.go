package routes

import (
	"Phylogeny/models"

	"github.com/gofiber/fiber/v2"
)

func NotFoundRoute(a *fiber.App) {
	a.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error:   "Not Found",
			Message: "The requested URL was not found on this server",
		})
	})
}
