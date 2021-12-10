package indexservice

import (
	"context"
	"github.com/evt/blockchain-api/internal/e"
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
func (svc *GroupService) GetIndex(ctx context.Context, id int64) (interface{}, e.Error) {
	group, err := svc.repo.GetIndex(ctx, id)
	if err != nil {
		return nil, e.NewInternalf("failed getting index from repository: %s", err)
	}

	return group, nil
}
