package blockhandler

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/evt/blockchain-api/internal/app/handlers/ctx"
	"github.com/gofiber/fiber/v2"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

// BlockHandler is a callback handler.
type BlockHandler struct {
	blockService BlockService
}

// New creates a new block handler.
func New(blockService BlockService) *BlockHandler {
	return &BlockHandler{
		blockService: blockService,
	}
}

// Get returns block by number or hash or "latest"
func (h *BlockHandler) Get(c *fiber.Ctx) error {
	blockIDStr := c.Params("id")
	var (
		block       map[string]interface{}
		blockNumber int64
		blockHash   common.Hash
		err         error
	)
	switch {
	case blockIDStr == "latest":
		block, err = h.blockService.GetBlockByNumber(c.Context(), nil)
		if err != nil {
			return ctx.Error(c, http.StatusInternalServerError, err)
		}
	case strings.HasPrefix(blockIDStr, "0x"):
		blockHash = common.HexToHash(blockIDStr)
		block, err = h.blockService.GetBlockByHash(c.Context(), blockHash)
		if err != nil {
			return ctx.Error(c, http.StatusInternalServerError, err)
		}
	default:
		blockNumber, err = strconv.ParseInt(blockIDStr, 10, 64)
		if err != nil {
			return ctx.Error(c, http.StatusBadRequest, err)
		}
		block, err = h.blockService.GetBlockByNumber(c.Context(), big.NewInt(blockNumber))
		if err != nil {
			return ctx.Error(c, http.StatusInternalServerError, err)
		}
	}

	return c.JSON(block)
}

// GetHeader returns block header by number or hash or "latest"
func (h *BlockHandler) GetHeader(c *fiber.Ctx) error {
	blockIDStr := c.Params("id")
	var (
		block       map[string]interface{}
		blockNumber int64
		blockHash   common.Hash
		err         error
	)
	switch {
	case blockIDStr == "latest":
		block, err = h.blockService.GetBlockHeaderByNumber(c.Context(), nil)
		if err != nil {
			return ctx.Error(c, http.StatusInternalServerError, err)
		}
	case strings.HasPrefix(blockIDStr, "0x"):
		blockHash = common.HexToHash(blockIDStr)
		block, err = h.blockService.GetBlockHeaderByHash(c.Context(), blockHash)
		if err != nil {
			return ctx.Error(c, http.StatusInternalServerError, err)
		}
	default:
		blockNumber, err = strconv.ParseInt(blockIDStr, 10, 64)
		if err != nil {
			return ctx.Error(c, http.StatusBadRequest, err)
		}
		block, err = h.blockService.GetBlockHeaderByNumber(c.Context(), big.NewInt(blockNumber))
		if err != nil {
			return ctx.Error(c, http.StatusInternalServerError, err)
		}
	}

	return c.JSON(block)
}
