package ctx

import (
	"github.com/evt/blockchain-api/internal/pkg/models"
	"github.com/gofiber/fiber/v2"
)

const unknownError = "unknown error"

// Error formats error before returning it in fiber response.
func Error(c *fiber.Ctx, status int, err error) error {
	if err == nil {
		return c.Status(status).JSON(models.Error{
			Error: unknownError,
		})
	}
	return c.Status(status).JSON(models.Error{
		Error: err.Error(),
	})
}
