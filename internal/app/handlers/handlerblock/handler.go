package handlerblock

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/evt/blockchain-api/internal/app/handlers/ctx"
	"github.com/evt/blockchain-api/internal/pkg/models"
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

	// swagger:route GET /blocks/:id blocks getBlock
	//
	// Returns block info from Ropsten.
	//
	// Returns block by ID which can be block number, block hash or "latest".
	// Block hash is identified as "0x...".
	//
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//       default: body:error
	//       200: body:block
	//       400: body:error

	blockIDStr := c.Params("id")
	var (
		block       *models.Block
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

	// swagger:route GET /blocks/:id/header blocks getBlockHeader
	//
	// Returns block info from Ropsten (header only).
	//
	// Returns block header by ID which can be block number, block hash or "latest".
	// Block hash is identified as "0x...".
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//       default: body:error
	//       200: body:blockNoBody
	//       400: body:error

	blockIDStr := c.Params("id")
	var (
		block       *models.Block
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

	blockNoBody := models.BlockNoBody{
		Header: block.Header,
	}

	return c.JSON(blockNoBody)
}
