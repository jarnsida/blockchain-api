package ctx

import (
	"github.com/gofiber/fiber/v2"
)

const unknownError = "unknown error"

// Error formats error before returning it in fiber response.
func Error(c *fiber.Ctx, status int, err error) error {
	if err == nil {
		return c.Status(status).JSON(fiber.Map{
			"error": unknownError,
		})
	}
	return c.Status(status).JSON(fiber.Map{
		"error": err.Error(),
	})
}
