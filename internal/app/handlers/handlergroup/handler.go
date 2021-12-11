package handlergroup

import (
	"github.com/evt/blockchain-api/internal/app/handlers/ctx"
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
		return ctx.Error(c, http.StatusInternalServerError, err)
	}

	return c.JSON(groups)
}

// Get handles GET /groups/:id request.
func (h *GroupHandler) Get(c *fiber.Ctx) error {

	// swagger:route GET /groups/:id groups getGroups
	//
	// Lists all contract groups.
	//
	// This will show all available contract groups.
	//
	//     Consumes:
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Deprecated: false
	//
	//     Security:
	//
	//     Responses:
	//       default: genericError
	//       200: group
	//       400: validationError

	groupIDStr := c.Params("id")
	groupID, err := strconv.ParseInt(groupIDStr, 10, 64)
	if err != nil {
		return ctx.Error(c, http.StatusBadRequest, err)
	}

	group, getGroupErr := h.groupService.GetGroup(c.Context(), groupID)
	if getGroupErr != nil {
		return ctx.Error(c, http.StatusInternalServerError, getGroupErr)
	}

	return c.JSON(group)
}
