package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func WebApiLogger(c *fiber.Ctx) error {
	// Record the start time
	start := time.Now()

	// Call the next middleware or handler
	err := c.Next()

	// Calculate the duration
	duration := time.Since(start)

	// Log the details of the request
	log.Printf(
		"[API Logger] %s %s %d - %v",
		c.Method(),
		c.OriginalURL(),
		c.Response().StatusCode(),
		duration,
	)

	return err
}
