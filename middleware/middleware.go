package middleware

import (
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func WebApiLogger(c *fiber.Ctx) error {
	if strings.HasSuffix(c.OriginalURL(), ".css") ||
		strings.HasSuffix(c.OriginalURL(), ".js") ||
		strings.HasSuffix(c.OriginalURL(), ".png") ||
		strings.HasSuffix(c.OriginalURL(), ".jpg") {
		return c.Next()
	}
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)

	log.Printf(
		"[API Logger] %s %s %d - %v",
		c.Method(),
		c.OriginalURL(),
		c.Response().StatusCode(),
		duration,
	)

	return err
}
