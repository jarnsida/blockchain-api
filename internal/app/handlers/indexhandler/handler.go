package indexhandler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

// IndexHandler is a callback handler.
type IndexHandler struct {
	indexService IndexService
}

// New creates a new index handler.
func New(indexService IndexService) *IndexHandler {
	return &IndexHandler{
		indexService: indexService,
	}
}

// Get handles GET /indexes/:id request.
func (h *IndexHandler) Get(c *fiber.Ctx) error {
	indexIDStr := c.Params("id")
	indexID, err := strconv.ParseInt(indexIDStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	index, getIndexErr := h.indexService.GetIndex(c.Context(), indexID)
	if err != nil {
		return c.Status(getIndexErr.Code()).JSON(fiber.Map{
			"error": getIndexErr.Detail(),
		})
	}

	return c.JSON(index)
}
