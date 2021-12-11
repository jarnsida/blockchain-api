package serviceblock

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evt/blockchain-api/internal/pkg/models"
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

// GetBlockByNumber returns block by number from repository.
func (svc *BlockService) GetBlockByNumber(ctx context.Context, number *big.Int) (*models.Block, error) {
	block, err := svc.repo.BlockByNumber(ctx, number)
	if err != nil {
		return nil, fmt.Errorf("failed getting block by number from repository: %w", err)
	}

	return &models.Block{
		Header: block.Header(),
		Body:   block.Body(),
	}, nil
}

// GetBlockHeaderByNumber returns block header by number from repository.
func (svc *BlockService) GetBlockHeaderByNumber(ctx context.Context, number *big.Int) (*models.Block, error) {
	header, err := svc.repo.HeaderByNumber(ctx, number)
	if err != nil {
		return nil, fmt.Errorf("failed getting block header by number from repository: %w", err)
	}

	return &models.Block{
		Header: header,
	}, nil
}

// GetBlockByHash returns block by hash from repository.
func (svc *BlockService) GetBlockByHash(ctx context.Context, hash common.Hash) (*models.Block, error) {
	block, err := svc.repo.BlockByHash(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("failed getting block by hash from repository: %w", err)
	}

	return &models.Block{
		Header: block.Header(),
		Body:   block.Body(),
	}, nil
}

// GetBlockHeaderByHash returns block by hash from repository.
func (svc *BlockService) GetBlockHeaderByHash(ctx context.Context, hash common.Hash) (*models.Block, error) {
	header, err := svc.repo.HeaderByHash(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("failed getting block header by hash from repository: %w", err)
	}

	return &models.Block{
		Header: header,
	}, nil
}
