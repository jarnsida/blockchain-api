//go:generate mockgen -destination=./mocks.go -source=./services.go -package=handlerblock

package handlerblock

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evt/blockchain-api/internal/pkg/models"
	"math/big"
)

// BlockService is a block service.
type BlockService interface {
	GetBlockByNumber(ctx context.Context, id *big.Int) (*models.Block, error)
	GetBlockByHash(ctx context.Context, hash common.Hash) (*models.Block, error)

	GetBlockHeaderByNumber(ctx context.Context, id *big.Int) (*models.Block, error)
	GetBlockHeaderByHash(ctx context.Context, hash common.Hash) (*models.Block, error)
}
