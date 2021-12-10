package grouphandler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

// GroupHandler is a callback handler.
type GroupHandler struct {
	groupService GroupService
}

// New creates a new group handler.
func New(groupService GroupService) *GroupHandler {
	return &GroupHandler{
		groupService: groupService,
	}
}

// GetAll handles GET /groups request.
func (h *GroupHandler) GetAll(c *fiber.Ctx) error {
	groups, err := h.groupService.GetGroupIDs(c.Context())
	if err != nil {
		return c.Status(err.Code()).JSON(fiber.Map{
			"error": err.Detail(),
		})
	}

	return c.JSON(groups)
}

// Get handles GET /group/:id request.
func (h *GroupHandler) Get(c *fiber.Ctx) error {
	groupIDStr := c.Params("id")
	groupID, err := strconv.ParseInt(groupIDStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	group, getGroupErr := h.groupService.GetGroup(c.Context(), groupID)
	if err != nil {
		return c.Status(getGroupErr.Code()).JSON(fiber.Map{
			"error": getGroupErr.Detail(),
		})
	}

	return c.JSON(group)
}
