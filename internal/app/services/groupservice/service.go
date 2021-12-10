package groupservice

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

// GetGroups returns all group IDs from repository.
func (svc *GroupService) GetGroupIDs(ctx context.Context) ([]int64, e.Error) {
	groupIDs, err := svc.repo.GetGroupIDs(ctx)
	if err != nil {
		return nil, e.NewInternalf("failed getting groups from repository: %s", err)
	}

	return groupIDs, nil
}

// GetGroup returns group from repository by ID.
func (svc *GroupService) GetGroup(ctx context.Context, id int64) (interface{}, e.Error) {
	group, err := svc.repo.GetGroup(ctx, id)
	if err != nil {
		return nil, e.NewInternalf("failed getting group from repository: %s", err)
	}

	return group, nil
}
