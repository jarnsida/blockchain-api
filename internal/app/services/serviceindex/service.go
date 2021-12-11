package serviceindex

import (
	"context"
	"fmt"
	"github.com/evt/blockchain-api/internal/pkg/models"
)

// GroupService is a group service.
type GroupService struct {
	repo Repository
}

// New creates a new group service.
func New(repo Repository) *GroupService {
	return &GroupService{
		repo: repo,
	}
}

// GetIndex returns index from repository by ID.
func (svc *GroupService) GetIndex(ctx context.Context, id int64) (*models.Index, error) {
	group, err := svc.repo.GetIndex(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed getting index from repository: %w", err)
	}

	return group, nil
}
