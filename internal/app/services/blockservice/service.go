package blockservice

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// BlockService is a block service.
type BlockService struct {
	repo Repository
}

// New creates a new block service.
func New(repo Repository) *BlockService {
	return &BlockService{
		repo: repo,
	}
}

// GetBlockByNumber returns block by number from repository by ID.
func (svc *BlockService) GetBlockByNumber(ctx context.Context, number int64) (map[string]interface{}, error) {
	block, err := svc.repo.BlockByNumber(ctx, big.NewInt(number))
	if err != nil {
		return nil, fmt.Errorf("failed getting block by number from repository: %w", err)
	}

	return map[string]interface{}{
		"header": block.Header(),
		"body":   block.Body(),
	}, nil
}

// GetBlockByHash returns block by hash from repository by ID.
func (svc *BlockService) GetBlockByHash(ctx context.Context, hash common.Hash) (map[string]interface{}, error) {
	block, err := svc.repo.BlockByHash(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("failed getting block by hash from repository: %w", err)
	}

	return map[string]interface{}{
		"header": block.Header(),
		"body":   block.Body(),
	}, nil
}
