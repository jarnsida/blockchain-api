//go:generate mockgen -destination=./mocks.go -source=./services.go -package=blockhandler

package blockhandler

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
)

// BlockService is a block service.
type BlockService interface {
	GetBlockByNumber(ctx context.Context, id int64) (map[string]interface{}, error)
	GetBlockByHash(ctx context.Context, hash common.Hash) (map[string]interface{}, error)
}
